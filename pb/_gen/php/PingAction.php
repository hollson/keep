<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: request.proto

/**
 * 客户端请求操作类型
 *
 * Protobuf type <code>PingAction</code>
 */
class PingAction
{
    /**
     * 心跳
     *
     * Generated from protobuf enum <code>Heartbeat = 0;</code>
     */
    const Heartbeat = 0;
    /**
     * 请求地图副本
     *
     * Generated from protobuf enum <code>ReqReplica = 1;</code>
     */
    const ReqReplica = 1;
    /**
     * 请求分发数据(副本/动作)
     *
     * Generated from protobuf enum <code>ReqDistribute = 2;</code>
     */
    const ReqDistribute = 2;
    /**
     * 通知下线
     *
     * Generated from protobuf enum <code>Offline = 3;</code>
     */
    const Offline = 3;

    private static $valueToName = [
        self::Heartbeat => 'Heartbeat',
        self::ReqReplica => 'ReqReplica',
        self::ReqDistribute => 'ReqDistribute',
        self::Offline => 'Offline',
    ];

    public static function name($value)
    {
        if (!isset(self::$valueToName[$value])) {
            throw new UnexpectedValueException(sprintf(
                    'Enum %s has no name defined for value %s', __CLASS__, $value));
        }
        return self::$valueToName[$value];
    }


    public static function value($name)
    {
        $const = __CLASS__ . '::' . strtoupper($name);
        if (!defined($const)) {
            throw new UnexpectedValueException(sprintf(
                    'Enum %s has no value defined for name %s', __CLASS__, $name));
        }
        return constant($const);
    }
}

