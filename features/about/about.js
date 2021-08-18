import {elementExist} from "/features/shared/util.js";

console.log('about.js loaded')
document.addEventListener("DOMContentLoaded", function () {
  if (elementExist("someDiv")) {
    GetSome();
    GetPartial();
  } else {
    console.log("element someDiv does not exist right now..");
  }
});
async function GetSome() {
  var response = await fetch("/about/api/somesome");
  var some = await response.json();
  console.log(some);
}
async function GetPartial() {
  var response = await fetch("/about/somep");
  var some = await response.text();
  const someDiv = document.getElementById("someDiv");
  someDiv.innerHTML = some;
  console.log(some);
}
