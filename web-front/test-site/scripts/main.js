let myImage = document.querySelector('img');

myImage.onclick = function() {
    let mySrc = myImage.getAttribute('src');
    if(mySrc === 'images/beautiful-girl.jpg') {
      myImage.setAttribute('src', 'images/girl2.jpg');
    } else {
      myImage.setAttribute('src', 'images/beautiful-girl.jpg');
    }
}
let myButton = document.querySelector('button');
let myHeading = document.querySelector('h1');

function setUserName() {
    let myName = prompt('请输入你的名字：');
    if(!myName || myName === null) {
      setUserName();
    } else {
      localStorage.setItem('name', myName);
      myHeading.textContent = '你很有眼光，' + myName;
    }
}

if(!localStorage.getItem('name')) {
    setUserName();
  } else {
    let storedName = localStorage.getItem('name');
    myHeading.textContent = '你很有眼光，' + storedName;
}

myButton.onclick = function() {
    setUserName();
}