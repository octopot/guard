<?php

if (strcasecmp(trim($_SERVER['REQUEST_URI'], '/'), 'protected') === 0) {
    phpinfo(INFO_ALL);
} else {
    echo '<h1>Welcome to Legacy project!</h1>', PHP_EOL;
}
