// click on letters and edit them in the big box


// color the pixels so that changing of colors in other functions
//   works without hiccups
function initPixels(color, className) {
    var pixels = document.getElementsByClassName(className);
    for (i = 0; i < pixels.length; i++) {
        pixels[i].style.backgroundColor = color;
    }
}

// Add event listeners to pixels
function addEventListeners(classname, miniLetterCallback) {
    var divs = document.querySelectorAll("." + classname);
    for (var i = 0; i < divs.length; i++) {
        divs[i].addEventListener("click", miniLetterCallback);
        divs[i].addEventListener("mouseenter", function(event) {
            if (event.target.style.backgroundColor != "rgb(252, 255, 245)"){
                event.target.style.backgroundColor = "rgb(62, 96, 111)";
            }
        });
        divs[i].addEventListener("mouseleave", function(event) {
            if (event.target.style.backgroundColor != "rgb(252, 255, 245)"){
                event.target.style.backgroundColor = "rgb(209, 219, 189)";
            }
        });
    }
}
// load small character to big character
var miniLetterCallback = function (event) {
    var newSelLet = event.toElement; // get Node that is getting clicked
    //console.log(event.toElement.className);
    if (event.toElement.className == "char_pix") { // if pixel gets clicked 
        newSelLet = event.toElement.parentNode.parentNode; // then get the parent
    }
    //console.log(newSelLet.className);
    var oldSelLet = selectedLetter;
    if (newSelLet === oldSelLet) { 
        return; // nothing to do
    } else {
        if (selectedLetter) {
            oldSelLet.style.backgroundColor = '#D1DBBD';
        }      
        newSelLet.style.backgroundColor = 'rgb(252, 255, 245)';
        //loadToBigChar(newSelLet);
	//copyLetter()
        copyLetter(newSelLet, document.querySelectorAll(".customc")[0]); // copy little letter to big
        selectedLetter = newSelLet; // 
    }
}

//   Adds events to pixels to enable color changing 
function addColorChanging() {
    var pixels = document.querySelectorAll(".c_pix");
    for (var i = 0; i < pixels.length; i++) {
        //pixels[i].style.backgroundColor = "#193441";
        pixels[i].addEventListener("click", function(event) {
            //console.log(event);
            stuff = event.toElement;
	         //console.log(stuff);
	         //          Needs to be in rgb to compare in javascript
            if (stuff.style.backgroundColor == "rgb(25, 52, 65)") { // #91AA9D
		          console.log("hello");
                stuff.style.backgroundColor = "rgb(145, 170, 157)";  //#193441
            } else {
                stuff.style.backgroundColor = "rgb(25, 52, 65)";
            }
            copyLetter(stuff.parentNode.parentNode, selectedLetter);
        });
    }
}

//   Copies one letter to another
//   Used to copy small letter to big letter
//     and big to small
function copyLetter(fromLetter, toLetter) {
    //console.log(fromLetter);
    //console.log(toLetter);
    var fromRows = fromLetter.childNodes;
    var toRows = toLetter.childNodes;
    //console.log(toRows[1]);
    for (var i = 0; i < fromRows.length; i++) {
        var fromPixels = fromRows[i].childNodes;
        var toPixels = toRows[i].childNodes;
        //console.log(fromPixels);
        for (var j = 0; j < fromPixels.length; j++) {
            toPixels[j].style.backgroundColor = fromPixels[j].style.backgroundColor;
            console.log(toPixels[j].style.backgroundColor + " to " +  fromPixels[j].style.backgroundColor);
        }
    }
}

function updateRangeValue(value, id) {
    //console.log(value);
    //console.log(id);
    var valueId = document.getElementById(id + "value");
    //console.log(valueId);
    valueId.innerHTML = value;
    //document.getElementById("valueId").innerHTML = value;
}

initPixels("rgb(145, 170, 157)", "c_pix");
initPixels("rgb(145, 170, 157)", "char_pix");
//var miniLetters = document.querySelectorAll(".char");
//console.log(miniLetters);
var selectedLetter = false;
addEventListeners("char", miniLetterCallback);
addColorChanging();
