function sendMail() {
  let name = document.getElementById("inputName").value;
  let phone = document.getElementById("inputPhone").value;
  let subject = document.getElementById("inputSubject").value;
  let message = document.getElementById("inputMessage").value;
  let email = document.getElementById("inputEmail")

  if (name == "") {
    return alert("Nama must be filled");
  } else if (email == "") {
    return alert("Email must be filled");
  } else if (phone == "") {
    return alert("Phone must be filled!");
  } else if (subject == "") {
    return alert("Subject must be choosen!");
  } else if (message == "") {
    return alert("Message must be filled!")
  }

  let emailReceiver = "chessarjunamariesto@gmail.com";

  let a = document.createElement("a");
  a.href = `mailto:${emailReceiver}?subject=${subject}&body=Halo, nama saya, ${name} ${message}. Silahkan kirimkan pesan saya di nomor ${phone}`;
  a.click();
}
