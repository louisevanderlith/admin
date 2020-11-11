import 'dart:html';
import 'package:mango_stock/bodies/category.dart';
import 'package:mango_ui/formstate.dart';
import 'package:mango_ui/keys.dart';
import 'package:mango_vehicle/bodies/engine.dart';

class VehicleEngine {

TextInputElement txtCode;
TextInputElement txtSerialNo;
NumberInputElement numOutput;
TextInputElement txtFuel;
TextInputElement txtDisplacement;

 VehicleEngine(){
   txtCode = querySelector("#txtCode");
   txtSerialNo = querySelector("#txtSerialNo");
   numOutput = querySelector("#numOutput");
   txtFuel = querySelector("#txtFuel");
   txtDisplacement = querySelector("#txtDisplacement");
  }

  String get code {
    return txtCode.value;
  }

  String get serialNo {
    return txtSerialNo.value;
  }

  num get output {
    return numOutput.valueAsNumber;
  }

  String get fuel {
    return txtFuel.value;
  }

  String get displacement {
    return txtDisplacement.value;
  }

  Engine toDto() {
    return new Engine(code, serialNo, output);
  }
}