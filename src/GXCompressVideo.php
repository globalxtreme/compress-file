<?php

namespace GlobalXtreme\CompressFile;

class GXCompressVideo
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
     * @param $filePath
     * @param $destination
     * @param $filename
     * @param $maxSize
     */
    public function __construct($filePath, $destination, $filename, $maxSize = 1000)
    {
        $this->filePath = $filePath;
        $this->destination = $destination;
        $this->filename = $filename;
        $this->maxSize = $maxSize;
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
     * @return false|object
     */
    public function process()
    {
        $params = json_encode([
            'path' => $this->filePath,
            'destination' => $this->destination,
            'filename' => $this->filename,
            'maxSize' => $this->maxSize
        ]);

        $result = exec(__DIR__ . '/go/resize-video ' . escapeshellarg($params), $return);
        $result = json_decode($result, true);
        if (isset($result['success'])) {
            return (object)$result;
        }

        return false;
    }

}
