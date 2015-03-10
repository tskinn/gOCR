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
    if (newSelLet.isSameNode(oldSelLet)) { 
        return; // nothing to do
    } else {
        if (selectedLetter) {
            oldSelLet.style.backgroundColor = '#D1DBBD';
        }      
        newSelLet.style.backgroundColor = '#91AA9D';
        //loadToBigChar(newSelLet);
	//copyLetter()
        copyLetter(newSelLet, document.querySelectorAll(".customc")[0]); // copy little letter to big
        selectedLetter = newSelLet; // 
    }
}

// function loadToBigChar(miniLetter) {
//     //var littleRows = miniLetter.querySelectorAll(".char_row");

//     //console.log(miniLetter.querySelectorAll(".c_pix"));
//     var littleRows = miniLetter;//.childNodes;
//     console.log("mini:" + miniLetter);
//     var bigRows = document.querySelectorAll(".c_row");
//     //console.log(miniLetter);
//     for (var i = 0; i < littleRows.length; i++) {
//         var pixels = littleRows.querySelectorAll(".char_pix");
//         //console.log("pixels: " + pixels);
//         var bigPixels = bigRows.querySelectorAll(".c_pix");
//         for (var j = 0; j < pixels.length; j++) {
//             if (pixels[j].style.backgroundColor == "#9B4C00") {
//                 bigPixels[j].style.backgroundColor = "#9B4C00";
//             } else {
//                 bigPixels[j].style.backgroundColor = "black";
//             }
//         }
//     }
// }

//   Adds events to pixels to enable color changing 
function addColorChanging() {
    var rows = document.querySelectorAll(".c_pix");
    for (var i = 0; i < rows.length; i++) {
        //rows[i].style.backgroundColor = "#9B4C00";
        rows[i].addEventListener("click", function(event) {
            //console.log(event);
            stuff = event.toElement;
	    console.log(stuff.style.backgroundColor);
            if (stuff.style.backgroundColor == "#91AA9D") {
                stuff.style.backgroundColor = "#193441";   // this code should be good
            } else {
                stuff.style.backgroundColor = "#91AA9D";
            }
            copyLetter(stuff.parentNode.parentNode, selectedLetter);
        });
    }
}

//   Copies one letter to another
//   Used to copy small letter to big letter
//     and big to small
function copyLetter(fromLetter, toLetter) {
    console.log(fromLetter);
    console.log(toLetter);
    var fromRows = fromLetter.childNodes;
    var toRows = toLetter.childNodes;
    //console.log(toRows[1]);
    for (var i = 0; i < fromRows.length; i++) {
        var fromPixels = fromRows[i].childNodes;
        var toPixels = toRows[i].childNodes;
        console.log(fromPixels);
        for (var j = 0; j < fromPixels.length; j++) {
            toPixels[j].style.backgroundColor = fromPixels[j].style.backgroundColor;
            console.log(toPixels[j].style.backgroundColor + " to " +  fromPixels[j].style.backgroundColor);
        }
    }
}


initPixels("#193441", "c_pix");
initPixels("#193441", "char_pix");
//var miniLetters = document.querySelectorAll(".char");
//console.log(miniLetters);
var selectedLetter = false;
addEventListeners("char", miniLetterCallback);
addColorChanging();
