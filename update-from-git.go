/*
 * @description       : Get latest files from GIT and allows to skip update certains file
 * @version           : "1.0.0"
 * @creator           : Gordon Lim <honwei189@gmail.com>
 * @created           : 26/05/2020 13:28:48
 * @last modified     : 02/06/2020 21:04:40
 * @last modified by  : Gordon Lim <honwei189@gmail.com>
 */

package main

import (
	"bufio"
	"fmt"
	"libs/utilities"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
	"time"

	// . "github.com/logrusorgru/aurora"
	"github.com/gookit/color"
	"github.com/howeyc/gopass"
	"github.com/urfave/cli"
)

var args = []string{}
var branch = []string{}
var cliArgs = []string{}
var cwd string
var gitDir string
var gitcred string
var lastUpdate string
var lastCommit string
var lastUpdateDateTime string
var lastUpdateDateOnly string
var lastUpdateTimeOnly string
var save string
var wd string
var today string
var username string
var password string
var startdate string
var enddate string

func main() {
	// var logcmd string
	var numFlags int

	// cmdb := "git"
	// args := "https://honwei189@github.com/honwei189/update-from-git.git"
	// cmd := exec.Command(cmdb, "pull", args)
	// stdin, e := cmd.StdinPipe()
	// if e != nil {
	// 	log.Fatal(e)
	// }

	// go func() {
	// 	defer stdin.Close()
	// 	io.WriteString(stdin, "QSRDFGHJfZERTYU")
	// }()

	// out, e := cmd.CombinedOutput()
	// if e != nil {
	// 	log.Fatal(string(out))
	// 	log.Fatal(e)
	// }

	// fmt.Printf("%s\n", out)

	// os.Exit(0)

	gitDir = ".git"
	cwd, _ = os.Getwd()
	wd = cwd
	currentTime := time.Now()
	today = fmt.Sprintf("%s", currentTime.Format("2006-01-02"))

	checkGitDir()

	app := cli.NewApp()
	app.Name = "update-from-git"
	app.EnableBashCompletion = true
	app.Usage = "\n\n\t\t\t" + color.FgLightCyan.Render("Get latest files from GIT and allows to skip update certains file")
	app.UsageText = color.FgRed.Render(app.Name + " command [command options] [arguments...]\n\n\t\t\texample:\n\n\t\t\t" + app.Name + "\n\n\t\t\t" + app.Name + " log\n\n\t\t\t" + app.Name + " log 2020-01-24\n\n\t\t\t" + app.Name + " log 2020-01-24 2020-04-02\n\n\t\t\t" + app.Name + " changed\n\n\t\t\t" + app.Name + " changed 2020-01-24\n\n\t\t\t" + app.Name + " changed 2020-01-24 2020-04-02\n\n\t\t\t" + app.Name + " --save=log.txt log\n\n\t\t\t" + app.Name + " --save=log.txt log 2020-01-24 2020-04-02\n\n\t\t\t" + app.Name + " -d .git2")
	app.Version = color.Yellow.Render("1.0.0")
	app.Compiled = time.Now()
	app.Authors = []cli.Author{
		cli.Author{
			Name:  color.FgLightCyan.Render("Gordon Lim"),
			Email: color.FgRed.Render("honwei189@gmail.com"),
		},
	}
	app.Copyright = color.FgMagenta.Render("2020 Gordon Lim") + "\n"

	// app.Flags = []cli.Flag{
	// 	cli.StringFlag{
	// 		Name:        "log, l",
	// 		Value:       "",
	// 		Usage:       color.FgLightGreen.Render("`Get changelog`"),
	// 		Destination: &logcmd,
	// 	},
	// }

	app.Action = func(c *cli.Context) error {
		numFlags = c.NumFlags()
		cliArgs = c.Args()

		if c.NumFlags() == 0 && len(cliArgs) == 0 {
			update()
			// changedFiles()
			// cli.ShowAppHelp(c)
			os.Exit(0)
		}

		return nil
	}

	app.UseShortOptionHandling = true
	app.Commands = []cli.Command{
		{
			Name:    "",
			Aliases: []string{""},
			Usage:   color.FgLightGreen.Render("Update files from GIT"),
			Action: func(c *cli.Context) error {
				// utilities.InitConf()
				color.Info.Println("Info message")
				return nil
			},
		},
		{
			Name:    "changed",
			Aliases: []string{""},
			Usage:   color.FgLightGreen.Render("Get changed files.  You can also passing specific date to get log from the date until to date"),
			Action: func(c *cli.Context) error {
				// utilities.InitConf()
				// color.Success.Println("Success message")

				// fmt.Printf("Hello %q", c.Args().Get(0))
				if len(c.Args()) >= 1 {
					if len(c.Args().Get(0)) > 0 {
						startdate = c.Args().Get(0)
						re := regexp.MustCompile("((19|20)\\d\\d)-(0?[1-9]|[12][0-9]|3[01])-(0?[1-9]|1[012])")

						// if re.Match(date) {

						// }

						if !re.MatchString(startdate) {
							fmt.Println()
							color.Error.Println("Invalid date format.  Date format must be YYYY-MM-DD  e.g: " + today)
							os.Exit(0)
						}
					}

					if len(c.Args()) > 1 && len(c.Args().Get(1)) > 0 {
						enddate = c.Args().Get(1)
						re := regexp.MustCompile("((19|20)\\d\\d)-(0?[1-9]|[12][0-9]|3[01])-(0?[1-9]|1[012])")

						// if re.Match(date) {

						// }

						if !re.MatchString(enddate) {
							fmt.Println()
							color.Error.Println("Invalid date format.  Date format must be YYYY-MM-DD  e.g: " + today)
							os.Exit(0)
						}
					}
				}

				// yourDate := convertDate(trimLastChar("Tue May 26 13:26:59 2020 +0800", 5))

				// fmt.Println(yourDate)

				// out, _ := utilities.CmdExec("git", "--git-dir=" + gitDir, "whatchanged", "-1")

				// color.Info.Println(out)
				changedFiles()
				return nil
			},
		},
		{
			Name:    "log",
			Aliases: []string{""},
			Usage:   color.FgLightGreen.Render("Get change log.  You can also passing specific date to get log from the date until to date"),
			Action: func(c *cli.Context) error {
				// utilities.InitConf()
				// color.Success.Println("Success message")

				// fmt.Printf("Hello %q", c.Args().Get(0))
				if len(c.Args()) >= 1 {
					if len(c.Args().Get(0)) > 0 {
						startdate = c.Args().Get(0)
						re := regexp.MustCompile("((19|20)\\d\\d)-(0?[1-9]|[12][0-9]|3[01])-(0?[1-9]|1[012])")

						// if re.Match(date) {

						// }

						if !re.MatchString(startdate) {
							fmt.Println()
							color.Error.Println("Invalid date format.  Date format must be YYYY-MM-DD  e.g: " + today)
							os.Exit(0)
						}
					}

					if len(c.Args()) > 1 && len(c.Args().Get(1)) > 0 {
						enddate = c.Args().Get(1)
						re := regexp.MustCompile("((19|20)\\d\\d)-(0?[1-9]|[12][0-9]|3[01])-(0?[1-9]|1[012])")

						// if re.Match(date) {

						// }

						if !re.MatchString(enddate) {
							fmt.Println()
							color.Error.Println("Invalid date format.  Date format must be YYYY-MM-DD  e.g: " + today)
							os.Exit(0)
						}
					}
				}

				// yourDate := convertDate(trimLastChar("Tue May 26 13:26:59 2020 +0800", 5))

				// fmt.Println(yourDate)

				// out, _ := utilities.CmdExec("git", "--git-dir=" + gitDir, "whatchanged", "-1")

				// color.Info.Println(out)

				changeLog()
				return nil
			},
		},
	}

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "dir, d",
			Value:       "",
			Usage:       color.FgLightGreen.Render("GIT repository directory. Default : .git"),
			Destination: &gitDir,
		},
		cli.StringFlag{
			Name:        "git, g",
			Value:       "",
			Usage:       color.FgLightGreen.Render("GIT repository directory. Default : .git"),
			Destination: &gitDir,
		},
		cli.StringFlag{
			Name:        "save, s",
			Value:       "",
			Usage:       color.FgLightGreen.Render("Save changelog or changed files list to specified file"),
			Destination: &save,
		},
		// cli.StringFlag{
		// 	Name:        "repo, r",
		// 	Value:       "",
		// 	Usage:       color.FgLightGreen.Render("GIT project directory"),
		// 	Destination: nil,
		// }
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

	if numFlags > 0 && len(cliArgs) == 0 {
		update()
		// changedFiles()
		// cli.ShowAppHelp(c)
		os.Exit(0)
	}
}

func changeLog() {
	_date := lastUpdateDateOnly
	_enddate := lastUpdateDateOnly

	if len(startdate) > 0 {
		_date = startdate
	}

	if len(enddate) > 0 {
		_enddate = enddate
	}

	fmt.Println()
	color.LightGreen.Print("Change log [ ")

	if startdate == enddate {
		color.Red.Print(_date)
	} else {
		color.LightGreen.Print("From ")
		color.Red.Print(_date)

		if len(enddate) > 0 {
			color.LightGreen.Print(" to ")
			color.Red.Print(_enddate)
		}
	}

	color.LightGreen.Println(" ]")
	fmt.Println("----------------------------------------------------------------------")

	logs, _ := utilities.CmdRun("git", "log", "--pretty=format:%B", "--abbrev-commit", "--date=relative", "origin", "master", "--since=\""+_date+" 12am\"", "--until=\""+_enddate+" 11:59pm\"")

	if len(logs) > 0 {
		i := 1
		writeLog := false
		logContent := ""

		if len(save) > 0 {
			writeLog = true

			if startdate == enddate {
				logContent += _date
			} else {
				logContent += "From " + _date

				if len(enddate) > 0 {
					logContent += " to " + _enddate
				}
			}

			logContent = fmt.Sprintf("Change log [ %s ]\r\n", logContent)
			logContent += fmt.Sprintf("----------------------------------------------------------------------\r\n")
		}

		for _, line := range logs {
			if strings.TrimSpace(line) != "" {
				fmt.Print(i)
				fmt.Println(". " + line)

				if writeLog {
					logContent += fmt.Sprintf("%d. %s\r\n", i, line)
				}
				i++
			}
		}

		if writeLog && len(logContent) > 0 {
			utilities.FilePutContents(save, logContent)
		}
		logContent = ""
	} else {
		fmt.Printf("\n\t\t\t%s\n", color.Red.Render("No record found"))
	}

	_date = ""
	_enddate = ""
	startdate = ""
	enddate = ""
}

func changedFiles() {
	_date := lastUpdateDateOnly
	_enddate := lastUpdateDateOnly

	if len(startdate) > 0 {
		_date = startdate
	}

	if len(enddate) > 0 {
		_enddate = enddate
	}

	fmt.Println()

	color.LightGreen.Print("Change log [ ")

	if startdate == enddate {
		color.Red.Print(_date)
	} else {
		color.LightGreen.Print("From ")
		color.Red.Print(_date)

		if len(enddate) > 0 {
			color.LightGreen.Print(" to ")
			color.Red.Print(_enddate)
		}
	}

	color.LightGreen.Println(" ]")
	// fmt.Println("=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=")
	fmt.Println("===============================================================================")
	fmt.Println()

	logs, _ := utilities.CmdRun("git", "log", "--pretty=format:%H", "--abbrev-commit", "--date=relative", "origin", "master", "--since=\""+_date+" 12am\"", "--until=\""+_enddate+" 11:59pm\"")

	i := 1
	highlight := 0
	code := make(map[string]string)
	code["A"] = "New"
	code["C"] = "Copy"
	code["D"] = "Delete"
	code["M"] = "Update" //Modified
	code["R"] = "Rename"
	code["U"] = "Update"
	code[" "] = "Unmodified"
	code["C75"] = "Copy"
	code["R097"] = "Move"
	code["R100"] = "Move"

	if len(logs) > 0 {
		writeLog := false
		logContent := ""

		if len(save) > 0 {
			writeLog = true

			if startdate == enddate {
				logContent += _date
			} else {
				logContent += "From " + _date

				if len(enddate) > 0 {
					logContent += " to " + _enddate
				}
			}

			logContent = fmt.Sprintf("Change log [ %s ]\r\n", logContent)
			logContent += fmt.Sprintf("----------------------------------------------------------------------\r\n")
		}

		for _, hash := range logs {
			logs, _ := utilities.CmdRun("git", "log", "-1", "--pretty=format:%cd %n%B%n-------------------------------------------------------------------------------", "--abbrev-commit", "--date=iso", "--name-status", hash)
			fmt.Print(i)
			fmt.Print(". ")
			// fmt.Println(". " + logs)

			if writeLog {
				logContent += fmt.Sprintf("%d.", i)
			}

			highlight = 0

			for _, line := range logs {
				if strings.TrimSpace(line) != "" {
					if highlight == 1 {
						// str := strings.Split(strings.TrimSpace(line), " ")
						// str1 := "Split   String on \nwhite    \tspaces."

						// re := regexp.MustCompile(`\S+`)

						// // fmt.Printf("Pattern: %v\n", re.String()) // Print Pattern

						// // fmt.Printf("String contains any match: %v\n", re.MatchString(line)) // True

						// submatchall := re.FindAllString(line, -1)
						// for _, element := range submatchall {
						// 	fmt.Println(element)
						// }

						str := utilities.RegSplit(line, `\S+`)

						if len(str) >= 1 {
							// color.BgBlue.Println(str[0])

							fmt.Printf("%s\t\t%s\n", color.Red.Render(code[strings.ToUpper(str[0])]), color.Success.Render(str[1]))

							if writeLog {
								logContent += fmt.Sprintf("%s\t\t%s\n", code[strings.ToUpper(str[0])], str[1])
							}
						}
					} else {
						if line == "-------------------------------------------------------------------------------" {
							highlight = 1
							color.LightMagenta.Println(line)
							if writeLog {
								logContent += fmt.Sprintf("%s\n", line)
							}
						} else {
							re := regexp.MustCompile("((19|20)\\d\\d)-(0?[1-9]|1[012])-(0?[1-9]|[12][0-9]|3[01]) ([0-9]{2}):([0-9]{2}):([0-9]{2}) [+-][0-9]{4}")

							if re.MatchString(line) {
								fmt.Println(utilities.ConvertUTCDateTime(strings.TrimSpace(line)))
								fmt.Println()
								if writeLog {
									logContent += fmt.Sprintf("%s\r\n\r\n", utilities.ConvertUTCDateTime(strings.TrimSpace(line)))
								}
							} else {
								// _l := utilities.SplitLines(line)

								// fmt.Println(_l)

								// if len(_l) == 1 {
								// 	fmt.Println(line)
								// } else {
								// 	for _, t := range _l {
								// 		fmt.Printf("- %s\n", strings.TrimSpace(t))
								// 	}
								// }

								// _l = nil

								fmt.Println(line)
								if writeLog {
									logContent += fmt.Sprintf("%s\r\n", line)
								}
							}
						}
					}
				}
			}
			fmt.Printf("\n\n\n\n")
			if writeLog {
				logContent += fmt.Sprintf("\n\n\n\n")
			}

			i++
		}

		if writeLog && len(logContent) > 0 {
			utilities.FilePutContents(save, logContent)
		}

		logContent = ""
	}

	_date = ""
	_enddate = ""
	startdate = ""
	enddate = ""
}

func changedDates() {
	// git log --graph --pretty=format:"%C(yellow)%h%x09%Creset%C(cyan)%C(bold)%ad%Creset  %C(green)%Creset %s" --date=short
	// git log --graph --pretty=format:"%C(cyan)%C(bold)%ad%Creset  %C(green)%Creset %s" --date=short
	// git log --graph --pretty=format:"%C(cyan)%C(bold)%cd %Creset  %C(green)%Creset %s" --date=short
}

func checkGitDir() {
	if len(gitDir) == 0 || strings.TrimSpace(gitDir) == "" {
		gitDir = ".git"
	}

	err := os.Chdir(wd)
	if err != nil {
		panic(err)
	}

	// fmt.Println("git ", gitDir)

	if utilities.DirExists(gitDir) && utilities.FileExists(gitDir+"/config") {
		_branch, _ := utilities.CmdExec("git", "--git-dir="+gitDir, "rev-parse", "--abbrev-ref", "@{u}")
		// branch = strings.Replace(branch, "/", " ", -1)
		branch = strings.Split(strings.TrimSpace(_branch), "/")
		lastUpdate, _ = utilities.CmdExec("git", "--git-dir="+gitDir, "rev-parse", "HEAD")
		lastUpdate = strings.TrimSpace(lastUpdate)

		// color.Info.Println(string(branch))
		// color.Info.Println(lastUpdate)

		_branch = ""

		// lastUpdateTimeOnly, _ = utilities.CmdExec("git", "--git-dir="+gitDir, "for-each-ref", "--format='%(committerdate)'", "--sort=-committerdate", "--count", "1")
		// lastUpdateTimeOnly = strings.Replace(lastUpdateTimeOnly, "'", "", -1)

		lastUpdateDateTime, _ = utilities.CmdExec("git", "--git-dir="+gitDir, "log", "-1", "--date=iso8601", "--pretty=format:%cd")
		_lastUpdate := strings.Split(strings.TrimSpace(lastUpdateDateTime), " ")
		lastUpdateDateOnly = _lastUpdate[0]
		lastUpdateTimeOnly = _lastUpdate[1]
		_lastUpdate = nil

		// lastUpdateDateOnly = utilities.ConvertUTCDateTime(lastUpdateTimeOnly)
		// lastUpdateTimeOnly = utilities.ConvertUTCTime(lastUpdateTimeOnly)

		// fmt.Printf("%q\n", re.Find([]byte(lastUpdateTimeOnly)))

		// matched, err := regexp.Match(`\D[+-][0-9]{4}\b`, []byte(lastUpdateTimeOnly))

		// fmt.Println(matched, err)

		// color.Info.Println(lastUpdateTimeOnly)

		// if runtime.GOOS == "windows" {

		// 	// lastCommit, _ := cmdRun("git", "--git-dir=" + gitDir, "ls-remote", strings.TrimSpace(branch))
		// 	// s, _ := cmdRun("date")
		// 	// printMemUsage()

		// 	// s, _ := cmdRun("git", "--git-dir=" + gitDir, "ls-files")

		// 	// for i := 0; i < len(s); i++ {
		// 	// 	color.Info.Println(s[i])
		// 	// }

		// 	// printMemUsage()

		// 	// s = nil

		// 	// Force GC to clear up, should see a memory drop
		// 	// runtime.GC()
		// 	// printMemUsage()
		// } else {
		// 	// utilities.CmdExec("git", "--git-dir="+gitDir, "config", "credential.helper", "store")
		// 	// utilities.CmdExec("git", "--git-dir="+gitDir, "config", "credential.helper", "'cache --timeout=600'") //Cache 10 minutes

		// 	// lastUpdateDateOnly, _ = utilities.CmdExec("echo", lastUpdateTimeOnly, "| sed 's/+.*$//g' | xargs -I{} date -d {} +\"%Y-%m-%d\"")
		// 	// lastUpdateTimeOnly, _ = utilities.CmdExec("echo", lastUpdateTimeOnly, "| sed 's/+.*$//g' | xargs -I{} date -d {} +\"%d/%m/%Y %H:%M:%S\"")
		// }

		// fmt.Println("YYYY-MM-DD ", fmt.Sprintf("%s", currentTime.Format("2006-01-02 03:04:05 PM")))
	} else {
		fmt.Println("")
		color.Error.Println("This is not a valid GIT repository directory")
		fmt.Println("")
		os.Exit(0)
	}
}

func update() {
	var _LastCommit string

	if len(gitDir) == 0 || strings.TrimSpace(gitDir) == "" {
		gitDir = ".git"
	}

	userLogin()

	fmt.Println()
	color.Error.Println("Checking repository...  Please wait... ")
	fmt.Println()

	if runtime.GOOS == "windows" {
		if len(branch) > 1 {
			_LastCommit, _ = utilities.CmdExec("git", "--git-dir="+gitDir, "ls-remote", branch[0], branch[1])
		} else {
			_LastCommit, _ = utilities.CmdExec("git", "--git-dir="+gitDir, "ls-remote", branch[0])
		}

		lastCommitFull := strings.Split(strings.TrimSpace(_LastCommit), "\t")
		lastCommit = lastCommitFull[0]

		branch = nil
		lastCommitFull = nil
		_LastCommit = ""
	} else {
		lastCommit, _ = utilities.CmdExec("git", "--git-dir="+gitDir, "ls-remote $(git rev-parse --abbrev-ref @{u} | sed 's/\\// /g') | cut -f1")
	}

	utilities.Clearscreen()

	if lastUpdate != lastCommit {
		color.Error.Println("Updating...  Please wait... ")
		fmt.Println()

		if utilities.FileExists(".ignores") {
			file, _ := os.Open(".ignores")
			defer file.Close()

			// if err != nil {
			// 	return err
			// }

			// Start reading from the file using a scanner.
			scanner := bufio.NewScanner(file)

			for scanner.Scan() {
				line := scanner.Text()
				if !strings.HasPrefix(line, "#") && len(line) > 0 {
					line = strings.TrimSpace(line)
					if strings.Contains(line, "/*") {
						err := filepath.Walk(strings.Replace(line, "/*", "", -1),
							func(path string, info os.FileInfo, err error) error {
								if err != nil {
									return err
								}

								if !info.IsDir() {
									// fmt.Println(path, info.Size())
									utilities.CmdExec("git", "--git-dir="+gitDir, "update-index", "--assume-unchanged", path)
								}

								return nil
							})
						if err != nil {
							log.Println(err)
						}
					} else {
						if utilities.Substr(line, -1) == "/" {
							_, err := os.Stat(utilities.Substr(line, 0, -1))

							if os.IsNotExist(err) {
							} else {
								utilities.CmdExec("git", "--git-dir="+gitDir, "update-index", "--assume-unchanged", line)
							}

							// utilities.CmdExec("git", "--git-dir=" + gitDir, "update-index", "--assume-unchanged", line)
						} else {
							// f, err := os.Open(line)

							// if err != nil {
							// 	// handle the error and return
							// 	// fmt.Println(err)
							// } else {
							// 	fi, err := f.Stat()
							// 	if err != nil {
							// 		// handle the error and return
							// 	}

							// 	if fi.IsDir() {
							// 		// it's a directory
							// 	} else {
							// 		fmt.Println(line + " is file")
							// 		// it's not a directory
							// 	}

							// 	fi = nil
							// }

							// defer f.Close()

							// f = nil

							fi, err := os.Stat(line)

							if os.IsNotExist(err) {
							} else {
								if fi.IsDir() {
									utilities.CmdExec("git", "--git-dir="+gitDir, "update-index", "--assume-unchanged", line+"/")
								} else {
									utilities.CmdExec("git", "--git-dir="+gitDir, "update-index", "--assume-unchanged", line)
								}
							}

							fi = nil
						}
					}
					// color.Info.Println(line)

					// color.Info.Println(err)
				}
			}

			if scanner.Err() != nil {
				fmt.Printf(" > Failed!: %v\n", scanner.Err())
			}

			_, e := utilities.CmdExec("git", "--git-dir="+gitDir, "pull")

			if e != nil {
				// log.Fatal(e)
				userReset()
				color.Error.Block(string(e.Error()))
				os.Exit(0)
			}

			changeLog()
			userReset()
		}
	} else {
		fmt.Println()
		// color.Error.Println("Already up-to-date.  Last update was " + lastUpdateDateTime)
		color.LightCyan.Print("Already up-to-date.  Last update was ")
		color.Red.Println(utilities.ConvertUTCDateTime(lastUpdateDateTime))
		fmt.Println()
		userReset()
	}
}

func userLogin() {
	utilities.Clearscreen()

	giturl, _ := utilities.CmdExec("git", "config", "--get", "remote.origin.url")

	gitcred, _ = utilities.CmdExec("git", "config", "--global", "credential.helper")
	gitlocalcred, _ := utilities.CmdExec("git", "config", "credential.helper")
	gitcred = strings.TrimSpace(gitcred)
	gitlocalcred = strings.TrimSpace(gitlocalcred)

	if gitlocalcred == "" || strings.Contains(gitlocalcred, "file") {
		_giturl := strings.Split(giturl, "/")

		color.BgRed.Println("Authentication for " + _giturl[0] + "//" + _giturl[2])
		fmt.Println()

		fmt.Print("Username: ")
		fmt.Scan(&username)

		// fmt.Println("")

		// fmt.Print("Enter fucking password: ")
		// fmt.Println("\033[8m") // Hide input
		// fmt.Scan(&password)
		// fmt.Println("\033[28m") // Show input

		// fmt.Printf("Enter silent password: ")
		gopass.GetPasswd() // Silent

		// input is in byte and need to convert to string
		// for storing and comparison
		// fmt.Println(string(silentPassword))

		fmt.Printf("Password: ")
		pass, _ := gopass.GetPasswdMasked() // Masked
		// fmt.Println(string(pass))

		password = string(pass)

		// giturl = strings.Replace(giturl, "https://", "https://"+username+":"+password+"@", 1)
		giturl = _giturl[0] + "//" + username + ":" + password + "@" + _giturl[2]

		if _giturl[0] == "http:" {
			giturl = fmt.Sprintf("%s\n%s", giturl, "https://"+username+":"+password+"@"+_giturl[2])
		}

		if runtime.GOOS == "windows" {
			utilities.FilePutContents(".git/creds", giturl)
			utilities.CmdExec("git", "config", "credential.helper", "store --file .git/creds")
		} else {
			utilities.CmdExec("git", "config", "credential.helper", "store")
			utilities.CmdExec("git", "config", "credential.helper", "cache --timeout 600") //Cache 10 minutes
		}

		fmt.Println()
		color.Cyan.Println("Logging in...  Please wait... ")

		_, err := utilities.CmdExec("git", "fetch", "--dry-run", "-q")

		// fmt.Println(err.Error())

		giturl = _giturl[0] + "//" + _giturl[2]
		_giturl = nil

		if err != nil {
			utilities.Clearscreen()
			color.Error.Println("Invalid username or password")
			fmt.Println()
			color.Red.Print("Authentication failed for ")
			color.Cyan.Println(giturl)

			userReset()
			os.Exit(1)
		}

		utilities.Clearscreen()
	}
}

func userReset() {
	utilities.CmdExec("git", "config", "credential.helper", gitcred)

	if runtime.GOOS == "windows" {
		utilities.Deletefile(".git/creds")
	}
}
