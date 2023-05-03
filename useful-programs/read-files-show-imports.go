
package main  
  
import (  
    "bufio"  
    "fmt"  
    "io/ioutil"  
    "os"  
    "path/filepath"  
    "regexp"  
    "strings"  
)  
  
func main() {  
    folder := `C:\Users\PabloZ\Desktop\CODE\GOLANG\fundamentals` // replace this with the path to your folder  
    outputFile := "./output.txt" // replace this with the path to your output file  
    
    fileInfos, err := ioutil.ReadDir(folder)  
    if err != nil {  
        fmt.Println(err)  
        return  
    }  
    
    // Compile a regular expression to match package import statements  
    importPattern := regexp.MustCompile(`import\s+["'`+"<](?P<PackagePath>[^\"'`>]+)[>'`\"]")  
    
    // Open the output file for writing  
    output, err := os.Create(outputFile)  
    if err != nil {  
        fmt.Println(err)  
        return  
    }  
    defer output.Close()  
    
    // Loop through all files in the folder and its subfolders  
    for _, fileInfo := range fileInfos {  
        filePath := filepath.Join(folder, fileInfo.Name())  
    
        if fileInfo.IsDir() {  
            // Recursively process subfolders  
            subFolder := filePath  
            main()  
            folderFiles, _ := ioutil.ReadDir(subFolder)  
            for _, folderFile := range folderFiles {  
                filePath := filepath.Join(subFolder, folderFile.Name())  
                processFile(output, filePath, importPattern)  
            }  
        } else {  
            // Process individual files  
            processFile(output, filePath, importPattern)  
        }  
    }  
}  
  
func processFile(output *os.File, filePath string, importPattern *regexp.Regexp) {  
    if strings.HasSuffix(filePath, ".go") {  
        // Open the file for reading  
        file, err := os.Open(filePath)  
        if err != nil {  
            fmt.Println(err)  
            return  
        }  
        defer file.Close()  
        
        // Scan the file line by line and extract package import statements  
        scanner := bufio.NewScanner(file)  
        for scanner.Scan() {  
            line := scanner.Text()  
            match := importPattern.FindStringSubmatch(line)  
            if len(match) > 1 {
                packageName := match[1]  
                output.WriteString(fmt.Sprintf("%s\t%s\n", filePath, packageName))  
            }
        }  
    }  
}
