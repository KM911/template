package util

import (
	"fmt"
	"os"
	"strings"
)

/*
返回文件夹下的 所有文件 和 一级文件夹
*/
func Dir(path string) ([]string, []string) {
	files, err := os.ReadDir(path)
	if err != nil {
		panic(err)
	}
	var fileNames []string
	var dirNames []string
	for _, file := range files {
		// 我们不希望就是显示文件夹 而是只显示文件
		if file.IsDir() {
			dirNames = append(dirNames, file.Name())
		} else {
			fileNames = append(fileNames, file.Name())
		}
	}
	return fileNames, dirNames
}

/*
判断是否为空目录
*/
func IsEmptyDir(path string) bool {
	filelist, folderlist := Dir(path)
	if len(filelist) == 0 && len(folderlist) == 0 {
		return true
	}
	return false
}
func SortByDepth(folders []string) []string {
	// Create a list of tuples, where each tuple contains the folder path and its depth.
	depths := make([]int, len(folders))
	for i, folder := range folders {
		// count depth
		depths[i] = strings.Count(folder, "/")
	}
	// Sort the list of tuples by the depth value.
	// from depth to shallower
	for i := 0; i < len(folders)-1; i++ {
		for j := 0; j < len(folders)-1-i; j++ {
			if depths[j] < depths[j+1] {
				depths[j], depths[j+1] = depths[j+1], depths[j]
				folders[j], folders[j+1] = folders[j+1], folders[j]
			}
		}
	}
	fmt.Println(folders)
	return folders
}

/*
返回文件夹下的所有 文件 和 文件夹
*/
func DeepDir(path string) ([]string, []string) {
	files, err := os.ReadDir(path)
	if err != nil {
		panic(err)
	}
	var fileNames []string
	var dirNames []string
	for _, file := range files {
		if file.IsDir() {
			dirNames = append(dirNames, file.Name())
			deepFileNames, deepDirNames := DeepDir(path + "/" + file.Name())
			for _, deepFileName := range deepFileNames {
				fileNames = append(fileNames, file.Name()+"/"+deepFileName)
			}
			for _, deepDirName := range deepDirNames {
				dirNames = append(dirNames, file.Name()+"/"+deepDirName)
			}
		} else {
			fileNames = append(fileNames, file.Name())
		}
	}
	return fileNames, dirNames
}

/*
返回文件夹下的所有 文件 和 文件夹
文件夹按深度排序
*/
func DeepDirSorted(path string) ([]string, []string) {
	files, err := os.ReadDir(path)
	if err != nil {
		panic(err)
	}
	var fileNames []string
	var dirNames []string
	for _, file := range files {
		if file.IsDir() {
			dirNames = append(dirNames, file.Name())
			deepFileNames, deepDirNames := DeepDir(path + "/" + file.Name())
			for _, deepFileName := range deepFileNames {
				fileNames = append(fileNames, file.Name()+"/"+deepFileName)
			}
			for _, deepDirName := range deepDirNames {
				dirNames = append(dirNames, file.Name()+"/"+deepDirName)
			}
		} else {
			fileNames = append(fileNames, file.Name())
		}
	}
	return fileNames, SortByDepth(dirNames)
}

/*
判定文件是否为目录
*/
func IsDir(file_ string) bool {
	fileInfo, err := os.Stat(file_)
	if err != nil {
		panic(err)
	}
	return fileInfo.IsDir()
}
