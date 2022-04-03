package main

func main() {
	//client := redis.NewClient(&redis.Options{
	//	Addr:     "127.0.0.1:6379", // dari server redis
	//	Password: "",
	//	DB:       0,
	//})
	//
	//pong, err := client.Ping().Result()
	//fmt.Println(pong, err)
	//errorSet := client.Set("name", "Gilang", 0).Err()
	//fmt.Println(errorSet)
	//stringGet, errGet := client.Get("name").Result()
	//fmt.Println(stringGet, errGet)

	Server().Run()
}

//
//func main() {
//	appConfig := config.NewConfig()
//	for {
//		var choice int
//		delivery.MainMenu()
//		fmt.Scan(&choice)
//		switch choice {
//		case 1:
//			delivery.MakeOrder(appConfig.UseCaseManager)
//		case 2:
//			delivery.FoodList(appConfig.UseCaseManager.FoodListUseCase())
//			utility.ChoiceToMainMenu()
//		case 3:
//			delivery.SearchFood(appConfig.UseCaseManager)
//		case 4:
//			delivery.Payment(appConfig.UseCaseManager)
//		case 5:
//			utility.SayGoodbay()
//		default:
//			fmt.Println("Tidak ada pilihan itu")
//		}
//	}
//}
