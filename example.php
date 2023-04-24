<?php
include('vendor/autoload.php');

$compressVideo = new \GlobalXtreme\CompressFile\GXCompressVideo('testing1.mp4', __DIR__ . "/", 'testing1.mp4');
if ($result = $compressVideo->process()) {
    echo $result->message;
    exit();
}

$compressImage = new \GlobalXtreme\CompressFile\GXCompressImage('testing1.jpg', __DIR__ . "/", 'testing1.jpg');
if ($result = $compressImage->process()) {
    echo $result->message;
    exit();
}

echo "error";
