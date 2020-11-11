import 'dart:html';

import 'package:mango_stock/bodies/category.dart';
import 'package:mango_ui/formstate.dart';
import 'package:mango_ui/keys.dart';
import 'package:mango_vehicle/bodies/gearbox.dart';

class VehicleGearbox {

TextInputElement txtCode;
TextInputElement txtSerialNo;
NumberInputElement numGears;
TextInputElement txtType;

 VehicleGearbox(){
   txtCode = querySelector("#txtCode");
   txtSerialNo = querySelector("#txtSerialNo");
   numGears = querySelector("#numGears");
   txtType = querySelector("#txtType");
  }

  String get code {
    return txtCode.value;
  }

  String get serialNo {
    return txtSerialNo.value;
  }

  num get gears {
    return numGears.valueAsNumber;
  }

  String get type {
    return txtType.value;
  }

  Gearbox toDto() {
    return new Gearbox(code, serialNo, gears, type);
  }
}