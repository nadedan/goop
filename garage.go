
package main

import (
  "fmt"
  "cars"
)

func main() {
  fmt.Println("hi")

  man_flex := cars.New_Car_Model(
    "Forge",
    "Flex",
    15000.23,
    "Blue",
    18.5,
    16,
  )
  man_flex.Set_Make("Ford")
  fmt.Println(man_flex)
  fmt.Println(man_flex.Range())

  flex := cars.New_Ford_Flex()
  fmt.Println(flex)

  wrangler := cars.New_Jeep_Wrangler()
  fmt.Println(wrangler)

  garage := []cars.Vehicle{flex, wrangler}
  fmt.Printf("\n\n My cars:\n")
  for _, car := range garage {
    fmt.Println(car.Make())
    fmt.Println(car.Model())
    fmt.Println(car.Price())
    fmt.Println(car.Range())

    crx, ok := car.(*cars.Honda_CRX)
    if ok {
      fmt.Println(crx.Little())
    }
  }

  fmt.Printf("\n Adding a car to my garage\n")
  garage = append(garage, cars.New_Honda_CRX())
  for _, car := range garage {
    fmt.Println(car.Make())
    fmt.Println(car.Model())
    fmt.Println(car.Price())
    fmt.Println(car.Range())

    crx, ok := car.(*cars.Honda_CRX)
    if ok {
      fmt.Println(crx.Little())
    }
  }
}
