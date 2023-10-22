# GoStrategyServer

Questo è il progetto del server in Golang per il gioco di strategia che vogliamo sviluppare

Allego anche lo schizzo della dinamica di connessione 
(NB: La struttura del progetto è stata costruita proprio su questo modello, quindi cercare di capire bene prima e poi scrivere codice!!)

[Server Architecture.pdf](https://github.com/GabryLinux/GoStrategyServer/files/13064115/Server.Architecture.pdf)
![Server Architecture](https://github.com/GabryLinux/GoStrategyServer/assets/87382811/19968095-1fcd-432c-9b8a-b0c68a141d66)
(PS: il .png è troppo piccolo, meglio scaricarsi e conservarsi il pdf, che è tuttavia totalmente modificabile, basta proporre e discutere eventuali cambiamenti)


Mancano ancora molte specifiche del progetto da decidere.
Abbozzo la lista di quello che ancora non è stato deciso, poi, in sede di riunione o in maniera totalmente asincrona (tipo su whatsapp) dare idee e tappare sti buchi:

<h1>Specifiche Server</h1>
<ul>
  <li>
    <h3>Linguaggio Di programmazione</h3>
    <ul>
      <p>Golang </p> 
      <a href="https://go.dev/doc/">link risorse</a>
    </ul>
  </li>
  <li>
    <h3>Protocollo Comunicazione Trasporto</h3>
    <ul>
      <p>TCP</p> 
      <a href="https://pkg.go.dev/net">link risorse</a>
    </ul>
  </li>
  <li>
    <h3>Protocollo interazione Client/Server</h3>
    <ul>
      <p>gRPC (o i tradizionali RPC forniti da Golang)</p> 
      <p><a href="https://grpc.io/">link risorse gRPC</a></p>
      <a href="https://pkg.go.dev/net/rpc">link risorse Golang RPC</a>
    </ul>
  </li>
  <li>
    <h3>Protocollo Autorizzazione e Autenticazione</h3>
    <ul>
      <p>?</p> 
      <a href="https://www.youtube.com/watch?v=dQw4w9WgXcQ">link risorse</a>
    </ul>
  </li>
   <li>
    <h3>Database</h3>
    <ul>
      <p>PostgreSQL</p> 
      <a href="https://blog.logrocket.com/building-simple-app-go-postgresql/">link risorse</a>
    </ul>
  </li>
  <li>
    <h3>Logica di Gioco</h3>
    <ul>
      <p>?</p> 
      <a href="https://www.youtube.com/watch?v=dQw4w9WgXcQ">link risorse</a>
    </ul>
  </li>
  <li>
    <h3>User State</h3>
    <ul>
      <p>?</p> 
      <a href="https://www.youtube.com/watch?v=dQw4w9WgXcQ">link risorse</a>
    </ul>
  </li>
  <li>
    <h3>Game State</h3>
    <ul>
      <p>?</p> 
      <a href="https://www.youtube.com/watch?v=dQw4w9WgXcQ">link risorse</a>
    </ul>
  </li>
  <li>
    <h3>Player State</h3>
    <ul>
      <p>?</p> 
      <a href="https://www.youtube.com/watch?v=dQw4w9WgXcQ">link risorse</a>
    </ul>
  </li>
  
  
</ul>


