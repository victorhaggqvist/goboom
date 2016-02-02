package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
	"sort"
	"strings"

	"github.com/go-ini/ini"
	"github.com/gocarina/gocsv"
	flag "github.com/ogier/pflag"
)

const version = 0

type Runnable struct {
	Cmd   string `csv:"name"`
	Count int    `csv:"count"`
}

type Config struct {
	DmenuParams string
	Ignore      []string `delim:","`
}

type CmdList []*Runnable

func (c CmdList) Len() int {
	return len(c)
}

func (c CmdList) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c CmdList) Less(i, j int) bool {
	return c[i].Count > c[j].Count
}

var (
	prePhase   bool
	postPhase  bool
	launcher   bool
	stats      bool
	dbFilePath string
	config     Config
)

func isInIgnoreList(item string) bool {
	for _, i := range config.Ignore {
		if i == item {
			return true
		}
	}
	return false
}

func addIfNotContains(items []string, item string) []string {
	for _, i := range items {
		if i == item {
			return items
		}
	}
	return append(items, item)
}

func openDB() map[string]int {
	if _, err := os.Stat(dbFilePath); err != nil {
		return make(map[string]int)
	}

	list := CmdList{}
	file, _ := os.Open(dbFilePath)
	if err := gocsv.UnmarshalFile(file, &list); err != nil {
		panic(err)
	}

	itemSet := make(map[string]int)
	for _, item := range list {
		itemSet[item.Cmd] = item.Count
	}
	return itemSet
}

func generatePath() []string {
	sysPath := os.Getenv("PATH")
	paths := strings.Split(sysPath, ":")

	pathItems := sort.StringSlice{}

	for _, p := range paths {
		if _, err := os.Stat(p); err == nil {
			files, _ := ioutil.ReadDir(p)
			for _, f := range files {
				if !isInIgnoreList(f.Name()) {
					pathItems = addIfNotContains(pathItems, f.Name())
				}
			}
		}
	}

	sort.Sort(pathItems)
	return pathItems
}

func rankPath(pathItems []string) CmdList {
	itemSet := openDB()

	rankedItems := CmdList{}
	for _, item := range pathItems {
		count, exists := itemSet[item]
		if !exists {
			count = 0
		}
		rankedItems = append(rankedItems, &Runnable{item, count})
	}

	sort.Sort(rankedItems)
	return rankedItems
}

func loadIni() {
	usr, _ := user.Current()
	configPath := filepath.Join(usr.HomeDir, ".goboom")
	iniFile := filepath.Join(configPath, "config.ini")
	dbFilePath = filepath.Join(configPath, "rankdb.csv")

	if _, err := os.Stat(configPath); err != nil {
		err := os.Mkdir(configPath, os.ModePerm)
		if err != nil {
			fmt.Println("failed to create config dir")
		}
	}

	config = Config{}
	if err := ini.MapToWithMapper(&config, ini.TitleUnderscore, iniFile); err != nil {
		panic(err)
	}
}

func updateRank(runnable string) {
	itemSet := openDB()

	count, exists := itemSet[runnable]
	if !exists {
		itemSet[runnable] = 1
	} else {
		itemSet[runnable] = count + 1
	}

	writeDB(itemSet)
}

func writeDB(itemSet map[string]int) {
	items := CmdList{}
	for runnable, count := range itemSet {
		items = append(items, &Runnable{runnable, count})
	}

	file, err := os.OpenFile(dbFilePath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	if err = gocsv.MarshalFile(&items, file); err != nil {
		panic(err)
	}
}

func displayStats() {
	items := openDB()

	sortedList := CmdList{}
	for runnable, count := range items {
		sortedList = append(sortedList, &Runnable{runnable, count})
	}

	sort.Sort(sortedList)

	fmt.Println("Name\tLaunch count")
	for _, i := range sortedList {
		fmt.Printf("%s\t%d\n", i.Cmd, i.Count)
	}
}

func main() {
	flag.BoolVar(&prePhase, "pre", false, "Generate dmenu in")
	flag.BoolVar(&launcher, "launcher", false, "Output launcher command")
	flag.BoolVar(&postPhase, "post", false, "Update ranking DB")
	flag.BoolVar(&stats, "stats", false, "View you statsB")
	flag.Parse()

	loadIni()

	if prePhase {
		pathList := generatePath()
		sortedList := rankPath(pathList)
		for _, item := range sortedList {
			fmt.Println(item.Cmd)
		}
	} else if postPhase {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()
		updateRank(input)
		fmt.Print(input)
	} else if launcher {
		fmt.Print("dmenu " + config.DmenuParams)

	} else if stats {
		displayStats()
	} else {
		fmt.Printf("goboom v%d (%s/%s/%s)\n", version, runtime.GOOS, runtime.GOARCH, runtime.Version())
		fmt.Println("\nTo actually use goboom execute goboom_run\n")
		flag.Usage()
	}
}
