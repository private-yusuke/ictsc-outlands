// @generated by protoc-gen-es v1.7.2 with parameter "target=ts,import_extension=none"
// @generated from file contestant/v1/problem.proto (package contestant.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from enum contestant.v1.QuestionType
 */
export enum QuestionType {
  /**
   * @generated from enum value: QUESTION_TYPE_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * @generated from enum value: QUESTION_TYPE_RADIO = 1;
   */
  RADIO = 1,

  /**
   * @generated from enum value: QUESTION_TYPE_CHECKBOX = 2;
   */
  CHECKBOX = 2,
}
// Retrieve enum metadata with: proto3.getEnumType(QuestionType)
proto3.util.setEnumType(QuestionType, "contestant.v1.QuestionType", [
  { no: 0, name: "QUESTION_TYPE_UNSPECIFIED" },
  { no: 1, name: "QUESTION_TYPE_RADIO" },
  { no: 2, name: "QUESTION_TYPE_CHECKBOX" },
]);

/**
 * @generated from enum contestant.v1.ProblemType
 */
export enum ProblemType {
  /**
   * @generated from enum value: PROBLEM_TYPE_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * @generated from enum value: PROBLEM_TYPE_DESCRIPTIVE = 1;
   */
  DESCRIPTIVE = 1,

  /**
   * @generated from enum value: PROBLEM_TYPE_MULTIPLE_CHOICE = 2;
   */
  MULTIPLE_CHOICE = 2,
}
// Retrieve enum metadata with: proto3.getEnumType(ProblemType)
proto3.util.setEnumType(ProblemType, "contestant.v1.ProblemType", [
  { no: 0, name: "PROBLEM_TYPE_UNSPECIFIED" },
  { no: 1, name: "PROBLEM_TYPE_DESCRIPTIVE" },
  { no: 2, name: "PROBLEM_TYPE_MULTIPLE_CHOICE" },
]);

/**
 * @generated from message contestant.v1.Choice
 */
export class Choice extends Message<Choice> {
  /**
   * @generated from field: int32 index = 1;
   */
  index = 0;

  /**
   * @generated from field: string body = 2;
   */
  body = "";

  constructor(data?: PartialMessage<Choice>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "contestant.v1.Choice";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "index", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "body", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Choice {
    return new Choice().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Choice {
    return new Choice().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Choice {
    return new Choice().fromJsonString(jsonString, options);
  }

  static equals(a: Choice | PlainMessage<Choice> | undefined, b: Choice | PlainMessage<Choice> | undefined): boolean {
    return proto3.util.equals(Choice, a, b);
  }
}

/**
 * @generated from message contestant.v1.Question
 */
export class Question extends Message<Question> {
  /**
   * @generated from field: string id = 1;
   */
  id = "";

  /**
   * @generated from field: string body = 2;
   */
  body = "";

  /**
   * @generated from field: contestant.v1.QuestionType type = 3;
   */
  type = QuestionType.UNSPECIFIED;

  /**
   * @generated from field: repeated contestant.v1.Choice choices = 4;
   */
  choices: Choice[] = [];

  /**
   * @generated from field: int32 point = 5;
   */
  point = 0;

  constructor(data?: PartialMessage<Question>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "contestant.v1.Question";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "body", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "type", kind: "enum", T: proto3.getEnumType(QuestionType) },
    { no: 4, name: "choices", kind: "message", T: Choice, repeated: true },
    { no: 5, name: "point", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Question {
    return new Question().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Question {
    return new Question().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Question {
    return new Question().fromJsonString(jsonString, options);
  }

  static equals(a: Question | PlainMessage<Question> | undefined, b: Question | PlainMessage<Question> | undefined): boolean {
    return proto3.util.equals(Question, a, b);
  }
}

/**
 * @generated from message contestant.v1.MultipleChoiceProblem
 */
export class MultipleChoiceProblem extends Message<MultipleChoiceProblem> {
  /**
   * @generated from field: optional string body = 1;
   */
  body?: string;

  /**
   * @generated from field: repeated contestant.v1.Question questions = 2;
   */
  questions: Question[] = [];

  constructor(data?: PartialMessage<MultipleChoiceProblem>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "contestant.v1.MultipleChoiceProblem";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "body", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
    { no: 2, name: "questions", kind: "message", T: Question, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MultipleChoiceProblem {
    return new MultipleChoiceProblem().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MultipleChoiceProblem {
    return new MultipleChoiceProblem().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MultipleChoiceProblem {
    return new MultipleChoiceProblem().fromJsonString(jsonString, options);
  }

  static equals(a: MultipleChoiceProblem | PlainMessage<MultipleChoiceProblem> | undefined, b: MultipleChoiceProblem | PlainMessage<MultipleChoiceProblem> | undefined): boolean {
    return proto3.util.equals(MultipleChoiceProblem, a, b);
  }
}

/**
 * @generated from message contestant.v1.ConnectionInfo
 */
export class ConnectionInfo extends Message<ConnectionInfo> {
  /**
   * @generated from field: string hostname = 1;
   */
  hostname = "";

  /**
   * @generated from field: string command = 2;
   */
  command = "";

  /**
   * @generated from field: string password = 3;
   */
  password = "";

  /**
   * @generated from field: string type = 4;
   */
  type = "";

  constructor(data?: PartialMessage<ConnectionInfo>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "contestant.v1.ConnectionInfo";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "hostname", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "command", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "password", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "type", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ConnectionInfo {
    return new ConnectionInfo().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ConnectionInfo {
    return new ConnectionInfo().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ConnectionInfo {
    return new ConnectionInfo().fromJsonString(jsonString, options);
  }

  static equals(a: ConnectionInfo | PlainMessage<ConnectionInfo> | undefined, b: ConnectionInfo | PlainMessage<ConnectionInfo> | undefined): boolean {
    return proto3.util.equals(ConnectionInfo, a, b);
  }
}

/**
 * @generated from message contestant.v1.DescriptiveProblem
 */
export class DescriptiveProblem extends Message<DescriptiveProblem> {
  /**
   * @generated from field: string body = 1;
   */
  body = "";

  /**
   * @generated from field: repeated contestant.v1.ConnectionInfo connection_infos = 2;
   */
  connectionInfos: ConnectionInfo[] = [];

  constructor(data?: PartialMessage<DescriptiveProblem>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "contestant.v1.DescriptiveProblem";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "body", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "connection_infos", kind: "message", T: ConnectionInfo, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DescriptiveProblem {
    return new DescriptiveProblem().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DescriptiveProblem {
    return new DescriptiveProblem().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DescriptiveProblem {
    return new DescriptiveProblem().fromJsonString(jsonString, options);
  }

  static equals(a: DescriptiveProblem | PlainMessage<DescriptiveProblem> | undefined, b: DescriptiveProblem | PlainMessage<DescriptiveProblem> | undefined): boolean {
    return proto3.util.equals(DescriptiveProblem, a, b);
  }
}

/**
 * @generated from message contestant.v1.Problem
 */
export class Problem extends Message<Problem> {
  /**
   * @generated from field: string id = 1;
   */
  id = "";

  /**
   * @generated from field: string title = 2;
   */
  title = "";

  /**
   * @generated from field: string code = 3;
   */
  code = "";

  /**
   * @generated from field: int32 point = 4;
   */
  point = 0;

  /**
   * @generated from field: contestant.v1.ProblemType type = 5;
   */
  type = ProblemType.UNSPECIFIED;

  /**
   * @generated from oneof contestant.v1.Problem.body
   */
  body: {
    /**
     * @generated from field: contestant.v1.DescriptiveProblem descriptive = 6;
     */
    value: DescriptiveProblem;
    case: "descriptive";
  } | {
    /**
     * @generated from field: contestant.v1.MultipleChoiceProblem multiple_choice = 7;
     */
    value: MultipleChoiceProblem;
    case: "multipleChoice";
  } | { case: undefined; value?: undefined } = { case: undefined };

  constructor(data?: PartialMessage<Problem>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "contestant.v1.Problem";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "title", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "code", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "point", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 5, name: "type", kind: "enum", T: proto3.getEnumType(ProblemType) },
    { no: 6, name: "descriptive", kind: "message", T: DescriptiveProblem, oneof: "body" },
    { no: 7, name: "multiple_choice", kind: "message", T: MultipleChoiceProblem, oneof: "body" },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Problem {
    return new Problem().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Problem {
    return new Problem().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Problem {
    return new Problem().fromJsonString(jsonString, options);
  }

  static equals(a: Problem | PlainMessage<Problem> | undefined, b: Problem | PlainMessage<Problem> | undefined): boolean {
    return proto3.util.equals(Problem, a, b);
  }
}

/**
 * @generated from message contestant.v1.GetProblemsRequest
 */
export class GetProblemsRequest extends Message<GetProblemsRequest> {
  constructor(data?: PartialMessage<GetProblemsRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "contestant.v1.GetProblemsRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetProblemsRequest {
    return new GetProblemsRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetProblemsRequest {
    return new GetProblemsRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetProblemsRequest {
    return new GetProblemsRequest().fromJsonString(jsonString, options);
  }

  static equals(a: GetProblemsRequest | PlainMessage<GetProblemsRequest> | undefined, b: GetProblemsRequest | PlainMessage<GetProblemsRequest> | undefined): boolean {
    return proto3.util.equals(GetProblemsRequest, a, b);
  }
}

/**
 * @generated from message contestant.v1.GetProblemsResponse
 */
export class GetProblemsResponse extends Message<GetProblemsResponse> {
  /**
   * @generated from field: repeated contestant.v1.Problem problems = 1;
   */
  problems: Problem[] = [];

  constructor(data?: PartialMessage<GetProblemsResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "contestant.v1.GetProblemsResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "problems", kind: "message", T: Problem, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetProblemsResponse {
    return new GetProblemsResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetProblemsResponse {
    return new GetProblemsResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetProblemsResponse {
    return new GetProblemsResponse().fromJsonString(jsonString, options);
  }

  static equals(a: GetProblemsResponse | PlainMessage<GetProblemsResponse> | undefined, b: GetProblemsResponse | PlainMessage<GetProblemsResponse> | undefined): boolean {
    return proto3.util.equals(GetProblemsResponse, a, b);
  }
}
