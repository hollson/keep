<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: protocol.proto

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * 相应结果
 *
 * Generated from protobuf message <code>PersonReply</code>
 */
class PersonReply extends \Google\Protobuf\Internal\Message
{
    /**
     *姓名
     *
     * Generated from protobuf field <code>string name = 1;</code>
     */
    private $name = '';
    /**
     *身份证
     *
     * Generated from protobuf field <code>string id_card = 2;</code>
     */
    private $id_card = '';
    /**
     *年龄
     *
     * Generated from protobuf field <code>int32 age = 3;</code>
     */
    private $age = 0;
    /**
     *性别
     *
     * Generated from protobuf field <code>.PersonReply.Sex sex = 4;</code>
     */
    private $sex = 0;
    /**
     *是否已婚
     *
     * Generated from protobuf field <code>bool married = 5;</code>
     */
    private $married = false;
    /**
     *资产
     *
     * Generated from protobuf field <code>double amount = 6;</code>
     */
    private $amount = 0.0;
    /**
     *地址
     *
     * Generated from protobuf field <code>repeated string address = 7;</code>
     */
    private $address;
    /**
     *其他
     *
     * Generated from protobuf field <code>repeated .google.protobuf.Any Others = 9;</code>
     */
    private $Others;
    /**
     * Generated from protobuf field <code>bytes data = 10;</code>
     */
    private $data = '';

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $name
     *          姓名
     *     @type string $id_card
     *          身份证
     *     @type int $age
     *          年龄
     *     @type int $sex
     *          性别
     *     @type bool $married
     *          是否已婚
     *     @type float $amount
     *          资产
     *     @type string[]|\Google\Protobuf\Internal\RepeatedField $address
     *          地址
     *     @type \Google\Protobuf\Any[]|\Google\Protobuf\Internal\RepeatedField $Others
     *          其他
     *     @type string $data
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Protocol::initOnce();
        parent::__construct($data);
    }

    /**
     *姓名
     *
     * Generated from protobuf field <code>string name = 1;</code>
     * @return string
     */
    public function getName()
    {
        return $this->name;
    }

    /**
     *姓名
     *
     * Generated from protobuf field <code>string name = 1;</code>
     * @param string $var
     * @return $this
     */
    public function setName($var)
    {
        GPBUtil::checkString($var, True);
        $this->name = $var;

        return $this;
    }

    /**
     *身份证
     *
     * Generated from protobuf field <code>string id_card = 2;</code>
     * @return string
     */
    public function getIdCard()
    {
        return $this->id_card;
    }

    /**
     *身份证
     *
     * Generated from protobuf field <code>string id_card = 2;</code>
     * @param string $var
     * @return $this
     */
    public function setIdCard($var)
    {
        GPBUtil::checkString($var, True);
        $this->id_card = $var;

        return $this;
    }

    /**
     *年龄
     *
     * Generated from protobuf field <code>int32 age = 3;</code>
     * @return int
     */
    public function getAge()
    {
        return $this->age;
    }

    /**
     *年龄
     *
     * Generated from protobuf field <code>int32 age = 3;</code>
     * @param int $var
     * @return $this
     */
    public function setAge($var)
    {
        GPBUtil::checkInt32($var);
        $this->age = $var;

        return $this;
    }

    /**
     *性别
     *
     * Generated from protobuf field <code>.PersonReply.Sex sex = 4;</code>
     * @return int
     */
    public function getSex()
    {
        return $this->sex;
    }

    /**
     *性别
     *
     * Generated from protobuf field <code>.PersonReply.Sex sex = 4;</code>
     * @param int $var
     * @return $this
     */
    public function setSex($var)
    {
        GPBUtil::checkEnum($var, \PersonReply_Sex::class);
        $this->sex = $var;

        return $this;
    }

    /**
     *是否已婚
     *
     * Generated from protobuf field <code>bool married = 5;</code>
     * @return bool
     */
    public function getMarried()
    {
        return $this->married;
    }

    /**
     *是否已婚
     *
     * Generated from protobuf field <code>bool married = 5;</code>
     * @param bool $var
     * @return $this
     */
    public function setMarried($var)
    {
        GPBUtil::checkBool($var);
        $this->married = $var;

        return $this;
    }

    /**
     *资产
     *
     * Generated from protobuf field <code>double amount = 6;</code>
     * @return float
     */
    public function getAmount()
    {
        return $this->amount;
    }

    /**
     *资产
     *
     * Generated from protobuf field <code>double amount = 6;</code>
     * @param float $var
     * @return $this
     */
    public function setAmount($var)
    {
        GPBUtil::checkDouble($var);
        $this->amount = $var;

        return $this;
    }

    /**
     *地址
     *
     * Generated from protobuf field <code>repeated string address = 7;</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getAddress()
    {
        return $this->address;
    }

    /**
     *地址
     *
     * Generated from protobuf field <code>repeated string address = 7;</code>
     * @param string[]|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setAddress($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::STRING);
        $this->address = $arr;

        return $this;
    }

    /**
     *其他
     *
     * Generated from protobuf field <code>repeated .google.protobuf.Any Others = 9;</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getOthers()
    {
        return $this->Others;
    }

    /**
     *其他
     *
     * Generated from protobuf field <code>repeated .google.protobuf.Any Others = 9;</code>
     * @param \Google\Protobuf\Any[]|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setOthers($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::MESSAGE, \Google\Protobuf\Any::class);
        $this->Others = $arr;

        return $this;
    }

    /**
     * Generated from protobuf field <code>bytes data = 10;</code>
     * @return string
     */
    public function getData()
    {
        return $this->data;
    }

    /**
     * Generated from protobuf field <code>bytes data = 10;</code>
     * @param string $var
     * @return $this
     */
    public function setData($var)
    {
        GPBUtil::checkString($var, False);
        $this->data = $var;

        return $this;
    }

}

