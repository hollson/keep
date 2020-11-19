// <auto-generated>
//     Generated by the protocol buffer compiler.  DO NOT EDIT!
//     source: protocol.proto
// </auto-generated>
#pragma warning disable 1591, 0612, 3021
#region Designer generated code

using pb = global::Google.Protobuf;
using pbc = global::Google.Protobuf.Collections;
using pbr = global::Google.Protobuf.Reflection;
using scg = global::System.Collections.Generic;
namespace Reworld.Teamwork {

  /// <summary>Holder for reflection information generated from protocol.proto</summary>
  public static partial class ProtocolReflection {

    #region Descriptor
    /// <summary>File descriptor for protocol.proto</summary>
    public static pbr::FileDescriptor Descriptor {
      get { return descriptor; }
    }
    private static pbr::FileDescriptor descriptor;

    static ProtocolReflection() {
      byte[] descriptorData = global::System.Convert.FromBase64String(
          string.Concat(
            "Cg5wcm90b2NvbC5wcm90bxohaW5jbHVkZS9nb29nbGUvcHJvdG9idWYvYW55",
            "LnByb3RvItsBCgtQZXJzb25SZXBseRIMCgRuYW1lGAEgASgJEg8KB2lkX2Nh",
            "cmQYAiABKAkSCwoDYWdlGAMgASgFEh0KA3NleBgEIAEoDjIQLlBlcnNvblJl",
            "cGx5LlNleBIPCgdtYXJyaWVkGAUgASgIEg4KBmFtb3VudBgGIAEoARIPCgdh",
            "ZGRyZXNzGAcgAygJEiQKBk90aGVycxgJIAMoCzIULmdvb2dsZS5wcm90b2J1",
            "Zi5BbnkSDAoEZGF0YRgKIAEoDCIbCgNTZXgSCgoGRkVNQUxFEAASCAoETUFM",
            "RRABQjQKD2NvbS5yZXdvcmxkLnd3d0IIVGVhbXdvcmtaBC47cGKqAhBSZXdv",
            "cmxkLlRlYW13b3JrYgZwcm90bzM="));
      descriptor = pbr::FileDescriptor.FromGeneratedCode(descriptorData,
          new pbr::FileDescriptor[] { global::Google.Protobuf.WellKnownTypes.AnyReflection.Descriptor, },
          new pbr::GeneratedClrTypeInfo(null, null, new pbr::GeneratedClrTypeInfo[] {
            new pbr::GeneratedClrTypeInfo(typeof(global::Reworld.Teamwork.PersonReply), global::Reworld.Teamwork.PersonReply.Parser, new[]{ "Name", "IdCard", "Age", "Sex", "Married", "Amount", "Address", "Others", "Data" }, null, new[]{ typeof(global::Reworld.Teamwork.PersonReply.Types.Sex) }, null, null)
          }));
    }
    #endregion

  }
  #region Messages
  /// <summary>
  /// 相应结果
  /// </summary>
  public sealed partial class PersonReply : pb::IMessage<PersonReply> {
    private static readonly pb::MessageParser<PersonReply> _parser = new pb::MessageParser<PersonReply>(() => new PersonReply());
    private pb::UnknownFieldSet _unknownFields;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pb::MessageParser<PersonReply> Parser { get { return _parser; } }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pbr::MessageDescriptor Descriptor {
      get { return global::Reworld.Teamwork.ProtocolReflection.Descriptor.MessageTypes[0]; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    pbr::MessageDescriptor pb::IMessage.Descriptor {
      get { return Descriptor; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public PersonReply() {
      OnConstruction();
    }

    partial void OnConstruction();

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public PersonReply(PersonReply other) : this() {
      name_ = other.name_;
      idCard_ = other.idCard_;
      age_ = other.age_;
      sex_ = other.sex_;
      married_ = other.married_;
      amount_ = other.amount_;
      address_ = other.address_.Clone();
      others_ = other.others_.Clone();
      data_ = other.data_;
      _unknownFields = pb::UnknownFieldSet.Clone(other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public PersonReply Clone() {
      return new PersonReply(this);
    }

    /// <summary>Field number for the "name" field.</summary>
    public const int NameFieldNumber = 1;
    private string name_ = "";
    /// <summary>
    ///姓名
    /// </summary>
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public string Name {
      get { return name_; }
      set {
        name_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
      }
    }

    /// <summary>Field number for the "id_card" field.</summary>
    public const int IdCardFieldNumber = 2;
    private string idCard_ = "";
    /// <summary>
    ///身份证
    /// </summary>
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public string IdCard {
      get { return idCard_; }
      set {
        idCard_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
      }
    }

    /// <summary>Field number for the "age" field.</summary>
    public const int AgeFieldNumber = 3;
    private int age_;
    /// <summary>
    ///年龄
    /// </summary>
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public int Age {
      get { return age_; }
      set {
        age_ = value;
      }
    }

    /// <summary>Field number for the "sex" field.</summary>
    public const int SexFieldNumber = 4;
    private global::Reworld.Teamwork.PersonReply.Types.Sex sex_ = global::Reworld.Teamwork.PersonReply.Types.Sex.Female;
    /// <summary>
    ///性别
    /// </summary>
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public global::Reworld.Teamwork.PersonReply.Types.Sex Sex {
      get { return sex_; }
      set {
        sex_ = value;
      }
    }

    /// <summary>Field number for the "married" field.</summary>
    public const int MarriedFieldNumber = 5;
    private bool married_;
    /// <summary>
    ///是否已婚
    /// </summary>
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public bool Married {
      get { return married_; }
      set {
        married_ = value;
      }
    }

    /// <summary>Field number for the "amount" field.</summary>
    public const int AmountFieldNumber = 6;
    private double amount_;
    /// <summary>
    ///资产
    /// </summary>
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public double Amount {
      get { return amount_; }
      set {
        amount_ = value;
      }
    }

    /// <summary>Field number for the "address" field.</summary>
    public const int AddressFieldNumber = 7;
    private static readonly pb::FieldCodec<string> _repeated_address_codec
        = pb::FieldCodec.ForString(58);
    private readonly pbc::RepeatedField<string> address_ = new pbc::RepeatedField<string>();
    /// <summary>
    ///地址
    /// </summary>
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public pbc::RepeatedField<string> Address {
      get { return address_; }
    }

    /// <summary>Field number for the "Others" field.</summary>
    public const int OthersFieldNumber = 9;
    private static readonly pb::FieldCodec<global::Google.Protobuf.WellKnownTypes.Any> _repeated_others_codec
        = pb::FieldCodec.ForMessage(74, global::Google.Protobuf.WellKnownTypes.Any.Parser);
    private readonly pbc::RepeatedField<global::Google.Protobuf.WellKnownTypes.Any> others_ = new pbc::RepeatedField<global::Google.Protobuf.WellKnownTypes.Any>();
    /// <summary>
    ///其他
    /// </summary>
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public pbc::RepeatedField<global::Google.Protobuf.WellKnownTypes.Any> Others {
      get { return others_; }
    }

    /// <summary>Field number for the "data" field.</summary>
    public const int DataFieldNumber = 10;
    private pb::ByteString data_ = pb::ByteString.Empty;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public pb::ByteString Data {
      get { return data_; }
      set {
        data_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override bool Equals(object other) {
      return Equals(other as PersonReply);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public bool Equals(PersonReply other) {
      if (ReferenceEquals(other, null)) {
        return false;
      }
      if (ReferenceEquals(other, this)) {
        return true;
      }
      if (Name != other.Name) return false;
      if (IdCard != other.IdCard) return false;
      if (Age != other.Age) return false;
      if (Sex != other.Sex) return false;
      if (Married != other.Married) return false;
      if (!pbc::ProtobufEqualityComparers.BitwiseDoubleEqualityComparer.Equals(Amount, other.Amount)) return false;
      if(!address_.Equals(other.address_)) return false;
      if(!others_.Equals(other.others_)) return false;
      if (Data != other.Data) return false;
      return Equals(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override int GetHashCode() {
      int hash = 1;
      if (Name.Length != 0) hash ^= Name.GetHashCode();
      if (IdCard.Length != 0) hash ^= IdCard.GetHashCode();
      if (Age != 0) hash ^= Age.GetHashCode();
      if (Sex != global::Reworld.Teamwork.PersonReply.Types.Sex.Female) hash ^= Sex.GetHashCode();
      if (Married != false) hash ^= Married.GetHashCode();
      if (Amount != 0D) hash ^= pbc::ProtobufEqualityComparers.BitwiseDoubleEqualityComparer.GetHashCode(Amount);
      hash ^= address_.GetHashCode();
      hash ^= others_.GetHashCode();
      if (Data.Length != 0) hash ^= Data.GetHashCode();
      if (_unknownFields != null) {
        hash ^= _unknownFields.GetHashCode();
      }
      return hash;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override string ToString() {
      return pb::JsonFormatter.ToDiagnosticString(this);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void WriteTo(pb::CodedOutputStream output) {
      if (Name.Length != 0) {
        output.WriteRawTag(10);
        output.WriteString(Name);
      }
      if (IdCard.Length != 0) {
        output.WriteRawTag(18);
        output.WriteString(IdCard);
      }
      if (Age != 0) {
        output.WriteRawTag(24);
        output.WriteInt32(Age);
      }
      if (Sex != global::Reworld.Teamwork.PersonReply.Types.Sex.Female) {
        output.WriteRawTag(32);
        output.WriteEnum((int) Sex);
      }
      if (Married != false) {
        output.WriteRawTag(40);
        output.WriteBool(Married);
      }
      if (Amount != 0D) {
        output.WriteRawTag(49);
        output.WriteDouble(Amount);
      }
      address_.WriteTo(output, _repeated_address_codec);
      others_.WriteTo(output, _repeated_others_codec);
      if (Data.Length != 0) {
        output.WriteRawTag(82);
        output.WriteBytes(Data);
      }
      if (_unknownFields != null) {
        _unknownFields.WriteTo(output);
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public int CalculateSize() {
      int size = 0;
      if (Name.Length != 0) {
        size += 1 + pb::CodedOutputStream.ComputeStringSize(Name);
      }
      if (IdCard.Length != 0) {
        size += 1 + pb::CodedOutputStream.ComputeStringSize(IdCard);
      }
      if (Age != 0) {
        size += 1 + pb::CodedOutputStream.ComputeInt32Size(Age);
      }
      if (Sex != global::Reworld.Teamwork.PersonReply.Types.Sex.Female) {
        size += 1 + pb::CodedOutputStream.ComputeEnumSize((int) Sex);
      }
      if (Married != false) {
        size += 1 + 1;
      }
      if (Amount != 0D) {
        size += 1 + 8;
      }
      size += address_.CalculateSize(_repeated_address_codec);
      size += others_.CalculateSize(_repeated_others_codec);
      if (Data.Length != 0) {
        size += 1 + pb::CodedOutputStream.ComputeBytesSize(Data);
      }
      if (_unknownFields != null) {
        size += _unknownFields.CalculateSize();
      }
      return size;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(PersonReply other) {
      if (other == null) {
        return;
      }
      if (other.Name.Length != 0) {
        Name = other.Name;
      }
      if (other.IdCard.Length != 0) {
        IdCard = other.IdCard;
      }
      if (other.Age != 0) {
        Age = other.Age;
      }
      if (other.Sex != global::Reworld.Teamwork.PersonReply.Types.Sex.Female) {
        Sex = other.Sex;
      }
      if (other.Married != false) {
        Married = other.Married;
      }
      if (other.Amount != 0D) {
        Amount = other.Amount;
      }
      address_.Add(other.address_);
      others_.Add(other.others_);
      if (other.Data.Length != 0) {
        Data = other.Data;
      }
      _unknownFields = pb::UnknownFieldSet.MergeFrom(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(pb::CodedInputStream input) {
      uint tag;
      while ((tag = input.ReadTag()) != 0) {
        switch(tag) {
          default:
            _unknownFields = pb::UnknownFieldSet.MergeFieldFrom(_unknownFields, input);
            break;
          case 10: {
            Name = input.ReadString();
            break;
          }
          case 18: {
            IdCard = input.ReadString();
            break;
          }
          case 24: {
            Age = input.ReadInt32();
            break;
          }
          case 32: {
            Sex = (global::Reworld.Teamwork.PersonReply.Types.Sex) input.ReadEnum();
            break;
          }
          case 40: {
            Married = input.ReadBool();
            break;
          }
          case 49: {
            Amount = input.ReadDouble();
            break;
          }
          case 58: {
            address_.AddEntriesFrom(input, _repeated_address_codec);
            break;
          }
          case 74: {
            others_.AddEntriesFrom(input, _repeated_others_codec);
            break;
          }
          case 82: {
            Data = input.ReadBytes();
            break;
          }
        }
      }
    }

    #region Nested types
    /// <summary>Container for nested types declared in the PersonReply message type.</summary>
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static partial class Types {
      public enum Sex {
        [pbr::OriginalName("FEMALE")] Female = 0,
        [pbr::OriginalName("MALE")] Male = 1,
      }

    }
    #endregion

  }

  #endregion

}

#endregion Designer generated code
