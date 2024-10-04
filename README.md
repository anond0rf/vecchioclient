<a name="readme-top"></a>
[![English](https://img.shields.io/badge/lang-en-blue.svg)](README-en.md) [![Italiano](https://img.shields.io/badge/lang-it-blue.svg)](README.md) 
![License](https://img.shields.io/github/license/anond0rf/vecchioclient) [![GitHub Pre-Release](https://img.shields.io/github/v/release/anond0rf/vecchioclient?include_prereleases&label=pre-release)](https://github.com/anond0rf/vecchioclient/releases) [![Go Version](https://img.shields.io/github/go-mod/go-version/anond0rf/vecchioclient)](https://github.com/anond0rf/vecchioclient)
<br />
<div align="center">
  <a href="https://github.com/anond0rf/vecchioclient">
    <img src="logo.png" alt="Logo" width="80" height="80">
  </a>
<h3 align="center">VecchioClient</h3>
  <p align="center">
    <strong>VecchioClient</strong> è una libreria Go per postare su <a href="https://vecchiochan.com/">vecchiochan.com</a>
    <br />
    <br />
    <a href="#installazione"><strong>Inizia »</strong></a>
    <br />
    <br />
    <a href="https://github.com/anond0rf/vecchioclient/blob/main/cmd/example-client/main.go">Guarda esempi</a>
    ·
    <a href="https://github.com/anond0rf/vecchioclient/issues">Segnala Bug</a>
    ·
    <a href="https://github.com/anond0rf/vecchioclient/issues">Richiedi Feature</a>
  </p>
</div>
 
## Caratteristiche

La libreria è basata sul reverse-engineering dell'endpoint `/post.php` di [NPFchan](https://github.com/fallenPineapple/NPFchan) ed espone un client che astrae i dettagli dell'invio del form e della gestione delle richieste http.  
Il client fornisce queste funzionalità:

- Pubblicare nuovi thread su board specifiche
- Rispondere a thread esistenti

È inoltre possibile una configurazione personalizzata che consente di passare `http.Client`, `User-Agent` e logger personalizzati, nonché di abilitare il logging dettagliato.  
Nessuna funzionalità di lettura viene fornita poiché NPFchan espone già l'[API](https://github.com/vichan-devel/vichan-API/) di vichan.

## Indice

1. [Installazione](#installazione)
2. [Utilizzo](#utilizzo)
   - [Utilizzo base del client](#utilizzo-base-del-client)
     - [Pubblicare un nuovo thread](#pubblicare-un-nuovo-thread)
     - [Pubblicare una risposta](#pubblicare-una-risposta)
   - [Configurazione personalizzata del client](#configurazione-personalizzata-del-client)
3. [Licenza](#licenza)

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
		log.Fatalf("Impossibile pubblicare il thread %+v. Errore: %v\n", thread, err)
	}
	fmt.Printf("Thread pubblicato con successo (id: %d) - %+v\n", id, thread)
    ```

    Se l'operazione va a buon fine, l'**id** del nuovo thread viene restituito.  
    NB: non è necessario impostare tutti i campi per istanziare la struct `Thread` e lo si può fare con un set ridotto:

    ```go
    thread := client.Thread{
		Board:    "b",
		Body:     "Questo è un nuovo thread sulla board /b/",   // Messaggio del thread
		Files:    []string{`C:\path\to\file.jpg`},
	}
    ```

    In questo caso, valori di default saranno assegnati agli altri campi.  
    **Board** è l'unico campo **obbligatorio** controllato dal client ma tieni presente che, poiché ogni board ha le sue impostazioni, potrebbero essere necessari più campi per postare (ad esempio, non è possibile postare un nuovo thread senza embed né file su /b/).

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
		log.Fatalf("Impossibile pubblicare la risposta %+v. Errore: %v\n", reply, err)
	}
	fmt.Printf("Risposta pubblicata con successo (id: %d) - %+v\n", id, reply)
    ```
    
    Se l'operazione va a buon fine, l'**id** della risposta viene restituito.  
    NB: non è necessario impostare tutti i campi per istanziare la struct `Reply` e lo si può fare con un set ridotto:

    ```go
    reply := client.Reply{
        Thread:   1,
		Board:    "b",
		Body:     "Questa è una nuova risposta al thread #1 della board /b/",   // Messaggio della risposta
	}
    ```

    In questo caso, valori predefiniti saranno assegnati agli altri campi.  
    **Board** e **Thread** sono gli unici campi **obbligatori** controllati dal client, ma tieni presente che, poiché ogni board ha le sue impostazioni, potrebbero essere necessari più campi per postare.

### Configurazione personalizzata del client

E' possibile passare una configurazione personalizzata al client creando una struct `Config` con i valori necessari come nell'esempio seguente:

```go
config := client.Config{
    Client:    &http.Client{Timeout: 10 * time.Second},                 // Client HTTP personalizzato
    Logger:    log.New(os.Stdout, "vecchioclient: ", log.LstdFlags),    // Logger personalizzato
    UserAgent: "MyCustomUserAgent/1.0",                                 // User-Agent personalizzato
    Verbose:   true,                                                    // Abilita/Disabilita il logging dettagliato
    
}

vc := client.NewVecchioClientWithConfig(config)
```

## Licenza

VecchioClient è concesso in licenza sotto la [Licenza LGPLv3](./LICENSE).

Questo significa che puoi usare, modificare e distribuire il software, a condizione che eventuali versioni modificate siano anch'esse concesse in licenza sotto la LGPLv3.

Per maggiori dettagli, consulta il testo completo della licenza nel file [LICENSE](./LICENSE).

Copyright © anond0rf
