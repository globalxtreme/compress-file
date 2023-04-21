<?php
include('vendor/autoload.php');

$compress = new \GlobalXtreme\CompressImage\GXCompressImage('testing.jpg', __DIR__ . "/", 'compressed_image.webp');
if ($result = $compress->process()) {
    echo $result->message;
    exit();
}

echo "error";
