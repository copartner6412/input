package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    // Open the input text file
    inputFile, err := os.Open("validate/tlds_alpha_by_domain.txt") // Replace with your input file name
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer inputFile.Close()

    // Create a map to store unique lowercase strings
    var tlds []string

    // Scan the file line by line
    scanner := bufio.NewScanner(inputFile)
    for scanner.Scan() {
        line := scanner.Text()
		words := strings.Fields(line)
        // Convert to lowercase and add to the map
        for _, word := range words {
            // Convert to lowercase and add to the map
            tld := strings.ToLower(word)
            tlds = append(tlds, tld)
    	}
	}

    // Check for errors during scanning
    if err := scanner.Err(); err != nil {
        fmt.Println("Error reading file:", err)
        return
    }

    // Create the output Go file
    outputFile, err := os.Create("tld_data.go") // Output file name
    if err != nil {
        fmt.Println("Error creating file:", err)
        return
    }
    defer outputFile.Close()

    // Write the package declaration and the map to the output file
    _, err = outputFile.WriteString("package main\n\n")
    if err != nil {
        fmt.Println("Error writing to file:", err)
        return
    }

    _, err = outputFile.WriteString("var TLDs = map[string]struct{}{\n")
    if err != nil {
        fmt.Println("Error writing to file:", err)
        return
    }

    // Write each unique lowercase string to the output file
    for _, str := range tlds {
        _, err = outputFile.WriteString(fmt.Sprintf("    \"%s\": {},\n", str))
        if err != nil {
            fmt.Println("Error writing to file:", err)
            return
        }
    }

    _, err = outputFile.WriteString("}\n")
    if err != nil {
        fmt.Println("Error writing to file:", err)
        return
    }

    fmt.Println("Output written to output.go")
}