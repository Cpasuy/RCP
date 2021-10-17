
<html>

<head>
    <meta charset="UTF-8"></meta>
    <meta http-equiv="X-UA-Compatible" content="IE=edge"></meta>
    <meta name="viewport" content="width=device-width, initial-scale=1.0"></meta>
    <title>Login page - RPC</title>
    <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.1.1/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-F3w7mX95PdgyTmZZMECAngseQB83DfGTowi0iMjiWaeVhAn4FJkqJByhZMI3AhiU" crossorigin="anonymous"></link>
</head>
<body>
    <header>
        <nav class="navbar navbar-expand-sm navbar-toggleable-sm navbar-light bg-white border-bottom box-shadow mb-3">
            <div class="container">
                <button class="navbar-toggler" type="button" data-toggle="collapse" data-target=".navbar-collapse" aria-controls="navbarSupportedContent"
                        aria-expanded="false" aria-label="Toggle navigation">
                    <span class="navbar-toggler-icon"></span>
                </button>
                <div class="navbar-collapse collapse d-sm-inline-flex justify-content-between">
                    <ul class="navbar-nav flex-grow-1">
                        <li class="nav-item">
                            <a class="nav-link text-dark" href="/rcp-auth">Home</a>
                        </li>
                        <li class="nav-item">
                            <a class="nav-link text-dark" href="#">Privacy</a>
                        </li>
                    </ul>
                </div>
            </div>
        </nav>
    </header>
    <br>
    </br>
    <div class="text-center">
      <div>
        <h1 class="text-center">Competitive Programming Network</h1> <br />
      <div>
      <a class="navbar-brand" href="#"><img src="https://pbs.twimg.com/profile_images/493847405670850561/qslkfHlq_400x400.jpeg" alt="RPC_Logo" width="100" height="100"/></a>
        <br></br>
      <div class="d-flex justify-content-center">
          <div class="row">
                <form action="" method="POST" >
                    <div class="form-group">
                        <label class="form-label" for="email">Email</label><input class="form-control" name="email" type="email" id="email" />
                    </div>
                    <div class="form-group">
                        <label class="form-label" for="username">Username</label><input class="form-control" name="username" type="text" id="username" />
                    </div>
                    <div class="form-group">
                        <label class="form-label" for="password">Password</label><input class="form-control" name="password" type="password" id="password" />
                    </div>
                    <div class="form-group">
                        <label class="form-label" for="confirmpwd">Confirm Password</label><input class="form-control" name="confirmpwd" type="password" id="confirmpwd" />
                    </div>
                    <div class="form-group">
                        <label class="form-label" for="firstname">First Name</label><input class="form-control" name="firstname" type="text" id="firstname" />
                    </div>
                    <div class="form-group">
                        <label class="form-label" for="lastname">Last Name</label><input class="form-control" name="lastname" type="text" id="lastname" />
                    </div>
                    <div class="form-group">
                        <label class="form-label" for="country">Country</label><input class="form-control" name="country" id="country"  />
                    </div>
                    <br></br>
                    <div class="btn-group" role="group">
                        <button class="btn btn-primary" type="submit" style="background: var(--bs-blue);">
                            Sign Up
                        </button>
                        
                    </div>
                </form>
            </div>
        </div>
    </div>
    </div>
    </div>
</body>
</html>

    