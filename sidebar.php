<?php
    echo '<html lang="en">
    <head>
        <meta charset="utf-8">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <meta http-equiv="X-UA-Compatible" content="IE=edge">
        <script src="https://ajax.googleapis.com/ajax/libs/jquery/3.3.1/jquery.min.js"></script>
    
        <title>Simply Fowl | Dispatch Page</title>
    
        <!-- Bootstrap CSS CDN -->
        <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css" integrity="sha384-ggOyR0iXCbMQv3Xipma34MD+dH/1fQ784/j6cY/iJTQUOhcWr7x9JvoRxT2MZw1T" crossorigin="anonymous">
        <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap-select@1.13.9/dist/css/bootstrap-select.min.css">
        <script src="https://cdn.jsdelivr.net/npm/bootstrap-select@1.13.9/dist/js/bootstrap-select.min.js"></script>
    
        <!-- Font Awesome JS -->
        <script src="https://kit.fontawesome.com/412b07d0f6.js" crossorigin="anonymous"></script><link rel="stylesheet" href="https://kit-free.fontawesome.com/releases/latest/css/free.min.css" media="all"><link rel="stylesheet" href="https://kit-free.fontawesome.com/releases/latest/css/free-v4-font-face.min.css" media="all"><link rel="stylesheet" href="https://kit-free.fontawesome.com/releases/latest/css/free-v4-shims.min.css" media="all">
        <script src="docCookies.js"></script>
    </head>
    <body data-gr-c-s-loaded="true">
        <!-- Sidebar  -->
        <div>
            <nav id="sidebar" class="fixed dispatch">
                <div class="sidebar-header"><h3 class="text-center">Simply Fowl</h3></div>
    
                <ul class="list-unstyled components">
                    <li><a href="saleshomepage.html"><i class="fas fa-home fa-fw" aria-hidden="true"></i> Home</a></li>
                    <li><a href="order.html"><i class="fas fa-box-open fa-fw" aria-hidden="true"></i> Orders</a></li>
                    <li class="active"><a href="#"><i class="fas fa-truck fa-fw" aria-hidden="true"></i> Dispatch</a></li>
                    <li><a href="shipping.html"><i class="fas fa-briefcase fa-fw" aria-hidden="true"></i> Shipments</a></li>
                    <li><a href="maintenance.html"><i class="fas fa-wrench fa-fw" aria-hidden="true"></i> Maintenance</a></li>
                    <li>
                        <a id="adminsub" href="#adminSubmenu" data-toggle="collapse" aria-expanded="false" class="dropdown-toggle">
                            <i class="fas fa-user-cog fa-fw" aria-hidden="true"></i> Admin</a>
                        <ul class="collapse list-unstyled" id="adminSubmenu">
                            <li><a href="FlockManager.html">Flock Manager</a></li>
                            <li><a href="accountsmanagement.html">Manage Accounts</a></li>
                        </ul>
                    </li>
            <li><a href="changePassword.html"><i class="fas fa-cog fa-fw"></i> Change Password </a></li>
                </ul>
    
            </nav>
        </div>
    </body>
    </html>'
?>