package main

/*
*
    WITHOUT CONTEXT
*
*/
func handler(w http.ResponseWriter, r *http.Request) {
    //create done channel to handle cancel routine
    doneCh := make(chan struct{}, 1)
    errCh := make(chan error, 1)
    go func() {
        errCh <- request(doneCh)
    }()

    // setting timeout use a seperate routine
    go func() {
        <-time.After(2 * time.Second)
        close(doneCh)
    }()

    select {
    case err := <-errCh:
        if err != nil {
            log.Println("failed:", err)
            return
        }
    }

    log.Println("success")
}

func request(doneCh chan struct{}) error {
    tr := &http.Transport{}
    client := &http.Client{Transport: tr}

    req, err := http.NewRequest("POST", backendService, nil)
    if err != nil {
        return err
    }
　　
    errCh := make(chan error, 1)
    go func() {
        _, err := client.Do(req)
        errCh <- err
    }()

    select {
    case err := <-errCh:
        if err != nil {
            return err
        }
    // if timeout thant cancel request
    case <-doneCh:
        tr.CancelRequest(req)
        <-errCh
        return fmt.Errorf("canceled")
    }

    return nil
}


/*
*
    WITH CONTEXT
*
*/
func handler(w http.ResponseWriter, r *http.Request) {
    // using context with timeout instead of create channel to handle timeout
    ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel()

    errCh := make(chan error, 1)
    go func() {
        errCh <- request(ctx)
    }()

    select {
    case err := <-errCh:
        if err != nil {
            log.Println("failed:", err)
            return
        }
    }

    log.Println("success")
}

func request(ctx context.Context) error {
    tr := &http.Transport{}
    client := &http.Client{Transport: tr}

    req, err := http.NewRequest("POST", backendService, nil)
    if err != nil {
        return err
    }

    errCh := make(chan error, 1)
    go func() {
        _, err := client.Do(req)
        errCh <- err
    }()

    select {
    case err := <-errCh:
        if err != nil {
            return err
        }
    case <-ctx.Done():
        tr.CancelRequest(req)
        <-errCh
        return ctx.Err()
    }

    return nil
}
