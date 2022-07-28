var myModal
var selected
var home="https://anotalo.herokuapp.com"  // produccion
//var home="http://localhost:3000"  // desarrollo
function buscar_notas() {
  document.getElementById("notas").innerHTML = ""
  function exito() {
    var datos = JSON.parse(this.responseText); //convertir a JSON
    html = "<ul>"
    for (x of datos) {
      id = x.id
      title = (typeof x.title === 'undefined') ? "" : x.title
      body = (typeof x.body === 'undefined') ? "" : x.body
      title = (title.length > 18) ? title.substr(0, 18) : title
      body = (body.length > 35) ? body.substr(0, 35) : body
      html = html + ` <li>
            <a href="#" onClick="update_note1('${id}')">
              <h2><i><b>${title}</b></i></h2>
              <br>
              <p>${body}</p>
            </a>
          </li>
          `
    }
    html = html + "</ul>"
    document.getElementById("notas").innerHTML = html
  }

  // funcion para la llamada fallida
  function error(err) {
    console.log('Solicitud fallida', err); //los detalles en el objecto "err"
  }

  var xhr = new XMLHttpRequest(); //invocar nueva instancia de XMLHttpRequest
  xhr.onload = exito; // llamar a la funcion exito si exitosa
  xhr.onerror = error;  // llamar a la funcion error si fallida
  xhr.open('GET', `${home}/note`); // Abrir solicitud GET
  xhr.send(); // mandar la solicitud al vervidor.
}



function add_note() {
  myModal = new bootstrap.Modal(document.getElementById("exampleModal"), {});
  var myModalEl = document.getElementById('exampleModal')
  myModalEl.addEventListener('shown.bs.modal', function (event) {
    document.getElementById("title").focus()
  })
  document.getElementById("title").value = ""
  document.getElementById("body").value = ""
  document.getElementById("labelModal").textContent = "Create"
  document.getElementById("create").style.display = "block";
  document.getElementById("update").style.display = "none";
  myModal.show();

}



function create_note() {
  title = document.getElementById("title").value
  body = document.getElementById("body").value

  let _data = {
    title: title,
    body: body,
  }

  fetch(`${home}/note`, {
    method: "POST",
    body: JSON.stringify(_data),
    headers: { "Content-type": "application/json; charset=UTF-8" }
  })
    .then(response => buscar_notas())
    .catch(err => console.log(err))
  myModal.hide()
}

function update_note() {
  title = document.getElementById("title").value
  body = document.getElementById("body").value

  let _data = {
    id: selected,
    title: title,
    body: body,
  }

  url = `${home}/note`
  fetch(url, {
    method: "PUT",
    body: JSON.stringify(_data),
    headers: { "Content-type": "application/json; charset=UTF-8" }
  })
    .then(response => { myModal.hide(); buscar_notas(); })
    .catch(err => console.log(err))
}



function delete_note() {

  url = `${home}/note/${selected}`
  fetch(url, {
    method: "DELETE",
    headers: { "Content-type": "application/json; charset=UTF-8" }
  })
    .then(response => { myModal.hide(); buscar_notas(); })
    .catch(err => console.log(err))
}

function update_note1(id) {

  function exito() {
    var datos = JSON.parse(this.responseText); //convertir a JSON
    selected = datos.id

    myModal = new bootstrap.Modal(document.getElementById("exampleModal"), {});
    var myModalEl = document.getElementById('exampleModal')
    myModalEl.addEventListener('shown.bs.modal', function (event) {
      document.getElementById("body").focus()
    })
    
    title = (typeof x.title === 'undefined') ? "" : datos.title
    body = (typeof x.body === 'undefined') ? "" : datos.body
    document.getElementById("title").value = title
    document.getElementById("body").value = body
    document.getElementById("labelModal").textContent = "Update"
    document.getElementById("create").style.display = "none";
    document.getElementById("update").style.display = "block";
    myModal.show();

  }

  // funcion para la llamada fallida
  function error(err) {
    console.log('Solicitud fallida', err); //los detalles en el objecto "err"
  }

  var xhr = new XMLHttpRequest(); //invocar nueva instancia de XMLHttpRequest
  xhr.onload = exito; // llamar a la funcion exito si exitosa
  xhr.onerror = error;  // llamar a la funcion error si fallida
  url = `${home}/note/${id}`
  xhr.open('GET', url); // Abrir solicitud GET
  xhr.send(); // mandar la solicitud al vervidor.
}

