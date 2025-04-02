package main

func main() {
	r := func() string {
		println("Olá, mundo!")
		return "Olá, mundo!"
	}()

	println(r)
}
