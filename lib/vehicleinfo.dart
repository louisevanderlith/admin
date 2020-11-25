import 'dart:html';

import 'package:mango_ui/keys.dart';

class VehicleInfo {
  HiddenInputElement txtVINKey;
  TextInputElement txtFullVIN;
  TextInputElement txtColour;
  TextInputElement txtPaintNo;
  SelectElement cboBodyStyle;
  NumberInputElement numDoors;
  CheckboxInputElement chkSpare;
  CheckboxInputElement chkService;
  TextInputElement txtCondition;
  TextInputElement txtIssues;
  NumberInputElement numMileage;

  VehicleGearbox() {
    txtVINKey = querySelector("#txtVINKey");
    txtFullVIN = querySelector("#txtFullVIN");
    txtColour = querySelector("#txtColour");
    txtPaintNo = querySelector("#txtPaintNo");
    cboBodyStyle = querySelector("#cboBodyStyle");
    numDoors = querySelector("#numDoors");
    chkSpare = querySelector("#chkSpare");
    chkService = querySelector("#chkService");
    txtCondition = querySelector("#txtCondition");
    txtIssues = querySelector("#txtIssues");
    numMileage = querySelector("#numMileage");
  }

  Key get vinKey {
    return new Key(txtVINKey.value);
  }

  String get fullVIN {
    return txtFullVIN.value;
  }

  String get colour {
    return txtColour.value;
  }

  String get paintNo {
    return txtPaintNo.value;
  }

  num get bodyStyle {
    return cboBodyStyle.selectedIndex;
  }

  num get doors {
    return numDoors.valueAsNumber;
  }

  bool get spare {
    return chkSpare.checked;
  }

  bool get service {
    return chkService.checked;
  }

  String get condition {
    return txtCondition.value;
  }

  String get issues {
    return txtIssues.value;
  }

  num get mileage {
    return numMileage.valueAsNumber;
  }
}
