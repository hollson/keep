<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: protocol.proto

namespace GPBMetadata;

class Protocol
{
    public static $is_initialized = false;

    public static function initOnce() {
        $pool = \Google\Protobuf\Internal\DescriptorPool::getGeneratedPool();

        if (static::$is_initialized == true) {
          return;
        }
        \GPBMetadata\GPBInclude\Google\Protobuf\Any::initOnce();
        $pool->internalAddGeneratedFile(hex2bin(
            "0aac020a0e70726f746f636f6c2e70726f746f22db010a0b506572736f6e" .
            "5265706c79120c0a046e616d65180120012809120f0a0769645f63617264" .
            "180220012809120b0a03616765180320012805121d0a0373657818042001" .
            "280e32102e506572736f6e5265706c792e536578120f0a076d6172726965" .
            "64180520012808120e0a06616d6f756e74180620012801120f0a07616464" .
            "7265737318072003280912240a064f746865727318092003280b32142e67" .
            "6f6f676c652e70726f746f6275662e416e79120c0a0464617461180a2001" .
            "280c221b0a03536578120a0a0646454d414c45100012080a044d414c4510" .
            "0142340a0f636f6d2e7265776f726c642e77777742085465616d776f726b" .
            "5a042e3b7062aa02105265776f726c642e5465616d776f726b620670726f" .
            "746f33"
        ), true);

        static::$is_initialized = true;
    }
}

