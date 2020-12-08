function loggedin() {
  //if no user is logged in and they are trying to access a page other than the login page,
  //send them to the login page
  if (
    docCookies.getItem("userid") == null &&
    window.location.href.indexOf("login.html") == -1
  ) {
    window.location.href = "login.html";
  }
}
