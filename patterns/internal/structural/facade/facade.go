// Package facade provides a simple example of facade pattern usage.
//
// The OpenWeatherMap API gives lots of information, so we are going to focus on getting live
// weather data in one city in some geo-located place by using its latitude and longitude
// values. The following are the requirements and acceptance criteria for this design pattern:
// 1. Provide a single type to access the data. All information retrieved from
// OpenWeatherMap service will pass through it.
// 2. Create a way to get the weather data for some city of some country.
// 3. Create a way to get the weather data for some latitude and longitude position.
// 4. Only second and thrird point must be visible outside of the package; everything
// else must be hidden (including all connection-related data).
package facade

type CurrentWeatherDataRetriever interface {
	GetByCityAndCountryCode(city, countryCode string) (Weather, error)
	GetByGeoCoordinates(lat, lon float32) (Weather, error)
}
