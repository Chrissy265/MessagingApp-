function loggedin() {
  //if no user is logged in and they are trying to access a page other than the login page,
  //send them to the login page

  alert(docCookies.getItem("userid"));

  /*if (
    docCookies.getItem("userid") == null &&
    window.location.href.indexOf("login.html") == -1
  ) {
    window.location.href = "login.html";
  }*/
}

function signOut() {
  docCookies.removeItem("email", email);
  docCookies.removeItem("clientID", clientID);
  docCookies.removeItem("userImage" + profile.getImageURL());
  docCookies.removeItem("token_uri", token_uri);
  docCookies.removeItem("clientSecret");
  docCookies.removeItem("userId", userId);
  var auth2 = gapi.auth2.getAuthInstance();
  auth2.signOut().then(function () {
    console.log("User signed out.");
  });
  window.location.href = "login.html";
}
