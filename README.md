GlobalXtreme Compress Image
======

[![Version](http://poser.pugx.org/globalxtreme/compress-image/version)](https://packagist.org/packages/globalxtreme/compress-image)
[![Total Downloads](http://poser.pugx.org/globalxtreme/compress-image/downloads)](https://packagist.org/packages/globalxtreme/compress-image)
[![License](http://poser.pugx.org/globalxtreme/compress-image/license)](https://packagist.org/packages/globalxtreme/compress-image)

### Install with composer

For install with [Composer](https://getcomposer.org/), simply require the
latest version of this package.

```bash
composer require globalxtreme/compress-image
```
##
#### Server
If you install this package in server (ubuntu or centos), you must be add script in **composer.json**.
```json
{
  "scripts": {
    "post-autoload-dump": [
      "Illuminate\\Foundation\\ComposerScripts::postAutoloadDump", 
      "@php artisan package:discover --ansi",
      "sh vendor/globalxtreme/compress-image/src/go/install-moduls.sh"
    ]
  }
} 
```

You can manual run **sh vendor/globalxtreme/compress-image/src/go/install-moduls.sh** from command line if you don't want to add it in **composer.json**
```bash
sh vendor/globalxtreme/compress-image/src/go/install-moduls.sh
```

##
#### Local
If you install this package in local (mac or windows or etc), you can temporarily disable or remove **"sh vendor/globalxtreme/compress-image/src/go/install-moduls.sh"** from **composer.json** and run this script.
```bash
cd vendor/globalxtreme/compress-image/src/go

go build resize.go
```