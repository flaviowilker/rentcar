<h3 align="center">
  RentCar
</h3>

<p align="center">A software to rent cars</p>

## About
This software is used as a system sample to rent cars. The software may help people to rent your cars.

## User Stories
* The customer can create an account to the system
* The customer can authenticate to the system
* The customer can rent a car
  * The customer can rent a car by one or more days
  * It's not possible to rent a car for an unavailable day or rented by another customer or himself
  * It's not possible to rent a car in the past
* The customer can cancel a rental
  * The customer cannot cancel a rental that happened
* The customer can see their rentals that will happen
* The administrator can register car models
* The administrator can register cars
* The administrator can register the days when the car can be rented
  * It's not possible to create another registration for the same day
  * It's not possible to create a registration in the past

## Architecture
 * This is an attempt to implement a clean architecture, that can be founded here https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html
