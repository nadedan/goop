
package cars

type Vehicle interface {
  Range() float32
  Make() string
  Model() string
  Price() float32
}

type car_base struct {
  color string
  mpg float32
  fuel_capacity float32
}

func (c car_base) Range() float32 {
  return c.mpg * c.fuel_capacity
}

type car_manufacturer struct {
  car_base
  manufacturer string
}

type Car_Model struct {
  car_manufacturer
  model string
  price float32
}


func (c car_manufacturer) Make() string {
  return c.manufacturer
}

func (c *car_manufacturer) Set_Make(manufacturer string) {
  c.manufacturer = manufacturer
}

func (c Car_Model) Model() string {
  return c.model
}

func (c *Car_Model) Set_Model(model string) {
  c.model = model
}

func (c Car_Model) Price() float32 {
  return c.price
}

func New_Car_Model(manufacturer  string,
                   model         string,
                   price         float32,
                   color         string,
                   mpg           float32,
                   fuel_capacity float32,
                  ) *Car_Model {
  c := new(Car_Model)
  c.manufacturer  = manufacturer
  c.model         = model
  c.price         = price
  c.color         = color
  c.mpg           = mpg
  c.fuel_capacity = fuel_capacity

  return c
}

type Ford_Flex struct {
  Car_Model
}

func New_Ford_Flex() *Ford_Flex {
  c := new(Ford_Flex)
  c.manufacturer  = "Ford"
  c.model         = "Flex"
  c.price         = 15000.23
  c.color         = "Blue"
  c.mpg           = 18.5
  c.fuel_capacity = 16

  return c
}

type Jeep_Wrangler struct {
  Car_Model
}

func New_Jeep_Wrangler() *Jeep_Wrangler {
  c := new(Jeep_Wrangler)
  c.manufacturer  = "Jeep"
  c.model         = "Wrangler"
  c.price         = 17920.89
  c.color         = "Black"
  c.mpg           = 22
  c.fuel_capacity = 18

  return c
}

type Honda_CRX struct {
  Car_Model
}

func New_Honda_CRX() *Honda_CRX {
  c := new(Honda_CRX)
  c.manufacturer  = "Honda"
  c.model         = "CRX"
  c.price         = 300.13
  c.color         = "Blue"
  c.mpg           = 30
  c.fuel_capacity = 12

  return c
}

