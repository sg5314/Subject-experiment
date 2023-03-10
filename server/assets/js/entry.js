const agent = window.navigator.userAgent.toLowerCase()


if (agent.indexOf('chrome') > -1) {
} else{
  if(!window.alert("Google Chromeで開いてください")){
    window.close();
  }
}


function clickEvent() {
  window.open('', '_self').close();
};

