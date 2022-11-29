let items = document.getElementsByTagName("li")
console.log("Welcome to Console Sir!")
for ( let i = 0; i < items.length; i++ ) {
    items[i].addEventListener("click", () => {
        items[i].classList.toggle("done")
    })
}