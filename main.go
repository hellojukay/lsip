package main

import "sync"

func main() {
	dirs := GetPathDirs()
	var wg sync.WaitGroup
	for index := range dirs {
		wg.Add(1)
		go func(dir string) {
			defer wg.Done()
			if files, err := ListDir(dir); err == nil {
				for _, file := range files {
					println(file)
				}
			}
		}(dirs[index])
	}
	wg.Wait()
}
