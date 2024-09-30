# VecchioClient

[![English](https://img.shields.io/badge/lang-en-blue.svg)](README-en.md) [![Italiano](https://img.shields.io/badge/lang-it-blue.svg)](README.md) 

**VecchioClient** è una libreria Go che fornisce un client per pubblicare su [vecchiochan.com](https://vecchiochan.com/).  
 
È basato sul reverse-engineering dell'endpoint `/post.php` di [NPFchan](https://github.com/fallenPineapple/NPFchan), astraendo i dettagli dell'invio del form e della gestione delle richieste http.

## Caratteristiche

- Pubblicare nuovi thread su board specifiche
- Rispondere a thread esistenti

È inoltre possibile una configurazione personalizzata che consente di passare `http.Client`, `User-Agent` e logger personalizzati, nonché di abilitare il logging dettagliato.

## Indice

1. [Installazione](#installazione)
2. [Utilizzo](#utilizzo)
   - [Utilizzo base del client](#utilizzo-base-del-client)
     - [Pubblicare un nuovo thread](#pubblicare-un-nuovo-thread)
     - [Pubblicare una risposta](#pubblicare-una-risposta)
   - [Configurazione personalizzata del client](#configurazione-personalizzata-del-client)

## Installazione

Per installare VecchioClient, usa `go get`:

```bash
go get github.com/anond0rf/vecchioclient
```

## Utilizzo

### Utilizzo base del client

VecchioClient offre un'API semplice e diretta per interagire con vecchiochan. Ecco come iniziare:

1. Importa il client nel tuo codice Go:

    ```go
    import "github.com/anond0rf/vecchioclient/client"
    ```

2. Crea un client:
   
    ```go
    vc := client.NewVecchioClient()
    ```

3. Usa il client per interagire con vecchiochan, ad esempio per pubblicare un nuovo thread o rispondere a uno esistente.

    - ##### Pubblicare un nuovo thread
    
    ```go
    thread := client.Thread{
		Board:    "b",
		Name:     "",
		Subject:  "",
		Email:    "",
		Spoiler:  false,
		Body:     "Questo è un nuovo thread sulla board /b/",   // Messaggio del thread
		Embed:    "",
		Password: "",
		Sage:     false,               // Impedisce il bumping e sostituisce l'email con "rabbia"
		Files:    []string{`C:\path\to\file.jpg`},
	}

    id, err := vc.NewThread(thread)
	if err != nil {
		log.Fatalf("Impossibile pubblicare il thread %+v. Errore: %v", thread, err)
	}
	fmt.Printf("Thread pubblicato con successo (id: %d) - %+v\n", id, thread)
    ```

    NB: non è necessario impostare tutti i campi per istanziare la struct `Thread` e lo si può fare con un set ridotto:

    ```go
    thread := client.Thread{
		Board:    "b",
		Body:     "Questo è un nuovo thread sulla board /b/",   // Messaggio del thread
		Files:    []string{`C:\path\to\file.jpg`},
	}
    ```

    In questo caso, valori di default saranno assegnati agli altri campi.  
    **Board** è l'unico campo **obbligatorio** controllato dal client ma tieni presente che, poiché le regole variano tra le board e ogni board ha le sue impostazioni, potrebbero essere necessari più campi per postare (ad esempio, non è possibile postare un nuovo thread senza embed né file su /b/).

    - ##### Pubblicare una risposta

    ```go
    reply := client.Reply{
		Thread:   1,
		Board:    "b",
		Name:     "",
		Email:    "",
		Spoiler:  false,
		Body:     "Questa è una nuova risposta al thread #1 della board /b/",    // Messaggio della risposta
		Embed:    "",
		Password: "",
		Sage:     false,            // Impedisce il bumping e sostituisce l'email con "rabbia"
		Files:    []string{`C:\path\to\file1.mp4`, `C:\path\to\file2.webm`},
	}

    id, err = vc.PostReply(reply)
	if err != nil {
		log.Fatalf("Impossibile pubblicare la risposta %+v. Errore: %v", reply, err)
	}
	fmt.Printf("Risposta pubblicata con successo (id: %d) - %+v\n", id, reply)
    ```

    NB: non è necessario impostare tutti i campi per istanziare la struct `Reply` e lo si può fare con un set ridotto:

    ```go
    reply := client.Reply{
        Thread:   1,
		Board:    "b",
		Body:     "Questa è una nuova risposta al thread #1 della board /b/",   // Messaggio della risposta
	}
    ```

    In questo caso, valori predefiniti saranno assegnati agli altri campi.  
    **Thread** è l'unico campo **obbligatorio** controllato dal client, ma tieni presente che, poiché le regole variano tra le board e ogni board ha le sue impostazioni, potrebbero essere necessari più campi per postare.

### Configurazione personalizzata del client

E' possibile passare una configurazione personalizzata al client creando una struct `Config` con i valori necessari come nell'esempio seguente:

```go
config := client.Config{
    Client:    &http.Client{Timeout: 10 * time.Second},                 // Client HTTP personalizzato
    Verbose:   true,                                                    // Abilita/Disabilita il logging dettagliato
    UserAgent: "MyCustomUserAgent/1.0",                                 // User-Agent personalizzato
    Logger:    log.New(os.Stdout, "vecchioclient: ", log.LstdFlags),    // Logger personalizzato
}

vc := client.NewVecchioClientWithConfig(config)
```