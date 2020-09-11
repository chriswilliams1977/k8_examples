//Reads data from a file
package datafile

import (
	"bufio"
	"log"
	"os"
)

func GetStrings(fileName string) ([]string,error){
	//declare slice to hold values
	var lines []string

	//open provided file
	file, err := os.Open(fileName)
	if err != nil{
		//if error opening file return array and error
		return nil, err
	}

	//returns a bufio.Scanner that reads from file
	scanner := bufio.NewScanner(file)
	//reads each line of file
	for scanner.Scan(){
		//Get line of text
		line := scanner.Text()
		//append text to slice
		lines = append(lines, line)
	}

	//close file to avoid consuming resources
	err = file.Close()
	if err != nil{
		//if error closing file return nil and error
		return nil, err
	}

	//checks to see if there was an issue with scanner while reading file
	if scanner.Err() != nil{
		////if error scanning file values return nil and error
		return nil, scanner.Err()
	}

	//if we got here we have slice populated and ready to return
	return lines, nil
}

func GetDataUsingSlice(filepath string)([]string, []int){
	//get lines from datafile
	lines, err := GetStrings(filepath)
	if err != nil{
		log.Fatal(err)
	}
	//create two slices, one to hold names and one to hold counts
	var names []string
	var counts []int
	//iterate over each line in file
	for _, line := range lines{
		matched := false
		//iterate over names slices
		for i, name := range names {
			//if name in file matches name in slice
			if name == line{
				//increment count slice
				counts[i]++
				matched = true
			}
		}
		//if no match
		if matched == false{
			//add name to name slice
			names = append(names, line)
			//set slice count to 1
			counts = append(counts, 1)
		}
	}
	return names, counts
}