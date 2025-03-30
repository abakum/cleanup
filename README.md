``` golang
	ctx, cancel := context.WithCancel(context.Background())
	defer closer.Close()
	cleanup := func() {
		log.Println("cleanup")
		<-ctx.Done()
		KidsDone(os.Getpid())
		log.Println("cleanup done" + DECTCEM + EL) // показать курсор, очистить строку
	}
	closer.Bind(cleanup)
	closer.Bind(cancel)

```
