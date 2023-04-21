<?php

namespace GlobalXtreme\CompressImage;

class GXCompressImage
{
    /**
     * @var string|null
     */
    protected $filePath;

    /**
     * @var string|null
     */
    protected $destination;

    /**
     * @var string|null
     */
    protected $filename;

    /**
     * @var int
     */
    protected $maxSize;

    /**
     * @var int
     */
    protected $quality;


    /**
     * @param $filePath
     * @param $destination
     * @param $filename
     * @param $maxSize
     * @param $quality
     */
    public function __construct($filePath, $destination, $filename, $maxSize = 1000, $quality = 70)
    {
        $this->filePath = $filePath;
        $this->destination = $destination;
        $this->filename = $filename;
        $this->maxSize = $maxSize;
        $this->quality = $quality;
    }

    /**
     * @param string|null $filePath
     *
     * @return GXCompressImage
     */
    public function setFilePath($filePath)
    {
        $this->filePath = $filePath;
        return $this;
    }

    /**
     * @param string|null $destination
     *
     * @return GXCompressImage
     */
    public function setDestination($destination)
    {
        $this->destination = $destination;
        return $this;
    }

    /**
     * @param string|null $filename
     *
     * @return GXCompressImage
     */
    public function setFilename($filename)
    {
        $this->filename = $filename;
        return $this;
    }

    /**
     * @param int $maxSize
     *
     * @return GXCompressImage
     */
    public function setMaxSize($maxSize)
    {
        $this->maxSize = $maxSize;
        return $this;
    }

    /**
     * @param int $quality
     *
     * @return GXCompressImage
     */
    public function setQuality($quality)
    {
        $this->quality = $quality;
        return $this;
    }


    /**
     * @return false|object
     */
    public function process()
    {
        $params = json_encode([
            'path' => $this->filePath,
            'destination' => $this->destination,
            'filename' => $this->filename,
            'maxSize' => $this->maxSize,
            'quality' => $this->quality
        ]);

        $result = exec(__DIR__ . '/go/resize-image ' . escapeshellarg($params), $return);
        $result = json_decode($result, true);
        if (isset($result['success'])) {
            return (object)$result;
        }

        return false;
    }

}
