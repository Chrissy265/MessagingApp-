import React from 'react';
import { GoogleLogout } from 'react-google-login';

const clientId =
  '554480865489-re2927613fjmo9hl5r96k9hf6tsicafo.apps.googleusercontent.com';

function Logout() {
  const onSuccess = () => {
    console.log('Logout made successfully');
    alert('Logout made successfully ✌');
  };

  return (
    <div>
      <GoogleLogout
        clientId={clientId}
        buttonText="Logout"
        onLogoutSuccess={onSuccess}
      ></GoogleLogout>
    </div>
  );
}

export default Logout;
