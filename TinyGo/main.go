package main

led := machine.LED

led.Configure(machine.PrinConfig{Mode: machine.PinOutput})

for{
	led.Low()
	time.Sleep(time.Millisecond * 500)

	led.High()
	time.Sleep(time.Millisecond * 500)
}