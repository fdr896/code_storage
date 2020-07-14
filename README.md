# Code storage
Web storage for your code snippets.


## Installation guide
  ### Installation using Docker
  - First of all download and install [Docker](https://www.docker.com/ "Motherfucking dooooocker") on your machine 
    >if you're using Windows do not choose `substitude Linux containers with Windows` option when installing the app
  - Launch Docker and check if it works (for example you can execute `docker run hello-world` in command prompt)
  - After that clone or download this project and open it in file manager (you supposed to see same files as in fdr896/code_storage)
  - Now everything depends on your OS
  
      Linux
    - run `build.sh` from console
    - wait till installation process finishs
    - run `run.sh` and wait several seconds
    - finally open http://localhost:5000 in browser and start using the application!
    
    Windows
    - tap on `build.bat` and wait till terminal close (probably some pop-up window will appear - in that case just give them what they want (access to your files or maybe some money:dollar::smiling_imp:)
    - tap on `run.bat`
    - wait when browser opens (if you see some erros instead of password page just reload your page several (not bigger than 7 millions:smile:) times
    - add your first code!
    
    Mac OS
    - use same approach as in Linux
  
  ### Manual installation
  - Install latest versions of [golang](https://golang.org/) and [npm](https://www.npmjs.com/) on your computer
  - Clone this repo and open it's main directory in terminal
  - `cd rest-api` and run `go run .` (some advertisments and server message will appear on your screen)
  - open new tab or new terminal and follow same steps as mentioned above but go to `svelte` directory instead of `rest-api`
  - type `npm install` and wait till all required node modules install
  - execute `npm run dev`
  - open http://localhost:5000 in browser
