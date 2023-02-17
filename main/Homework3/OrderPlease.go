package kata

import "strings"

//Task 5 https://www.codewars.com/kata/55c45be3b2079eccff00010f
func Order(sentence string) string {
  
  var numberFind []string
  var stringSplit []string 
  stringSplit = strings.Split(sentence, " ")
  stringOrder := make([]string, len(stringSplit))
  
  
  for i := 0; i < len(stringSplit); i++{
    
    numberFind = strings.Split(stringSplit[i], "")
    
    for m := 0; m < len(numberFind); m++{
      
      switch numberFind[m]{
        case "1":
        stringOrder[0] = stringSplit[i]
        break
        
        case "2":
        stringOrder[1] = stringSplit[i]
        break
        
        case "3":
        stringOrder[2] = stringSplit[i]
        break
        
        case "4":
        stringOrder[3] = stringSplit[i]
        break
        
        case "5":
        stringOrder[4] = stringSplit[i]
        break
      
        case "6":
        stringOrder[5] = stringSplit[i]
        break
        
        case "7":
        stringOrder[6] = stringSplit[i]
        break
        
        case "8":
        stringOrder[7] = stringSplit[i]
        break
        
        case "9":
        stringOrder[8] = stringSplit[i]
        break
        
      
      }
      
    }
    sentence = strings.Join(stringOrder," ")
  }
  
  
  
  return sentence
}