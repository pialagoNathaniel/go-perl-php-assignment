// Pialago, Nathaniel Christian M.  BSCS601
package main
import "fmt"

func main() {
    var choice string
    fmt.Println("Good Morning! Did you sleep well? [y/n]")
    fmt.Scanln(&choice)
    
    if choice == "y" {
        fmt.Println("That's nice! Now take 30 minutes of exercise!" + choice)
    } else {
        fmt.Println("That's okay! Take 15 minutes of meditation!")
    }
    
    fmt.Println("Once you're done, have a breakfast.")
    fmt.Println("Did you take a shower in the morning? [y/n]")
    fmt.Scanln(&choice)
    
    if choice != "y" {
        fmt.Println("Alright, take a shower now!")
    }
    fmt.Println("Get dressed!")
    fmt.Println("Do you feel ready to start the day? [y/n]")
    fmt.Scanln(&choice)
    
    if choice == "y" {
        fmt.Println("Cool! You should make a to-do list!")
    } else {
        fmt.Println("It's okay, take 15 minutes of self-care or self-love!")
    }
    
    fmt.Println("Are you feeling better now? [y/n]")
    fmt.Scanln(&choice)
    
    if choice == "y" {
        fmt.Println("You are now ready for a strong day!")
    } else {
        fmt.Println("I see... Take a test!")
    }
}