import 'dart:html';

import 'package:dart_toast/dart_toast.dart';
import 'package:mango_ui/formstate.dart';
import 'package:mango_ui/keys.dart';
import 'package:mango_wear/bodies/clothing.dart';
import 'package:mango_wear/wearapi.dart';

class ClothingForm extends FormState {
  Key objKey;
  TextInputElement txtCode;
  TextInputElement txtDescription;
  SelectElement cboBrand;
  SelectElement cboType;
  SelectElement cboSize;
  TextInputElement txtColour;
  TextInputElement txtMaterial;
  NumberInputElement numWeight;

  ClothingForm(Key k) : super("#frmClothing", "#btnSubmit") {
    objKey = k;
    txtCode = querySelector("#txtCode");
    txtDescription = querySelector("#txtDescription");
    cboBrand = querySelector("#cboBrand");
    cboType = querySelector("#cboType");
    cboSize = querySelector("#cboSize");
    txtColour = querySelector("#txtColour");
    txtMaterial = querySelector("#txtMaterial");
    numWeight = querySelector("#numWeight");

    cboType.onChange.listen(onTypeChange);

    querySelector("#btnSubmit").onClick.listen(onSend);
  }

  void onTypeChange(Event e) {
    //update sizes
  }

  String get code {
    return txtCode.value;
  }

  Key get brand {
    return new Key(cboBrand.value);
  }

  String get description {
    return txtDescription.value;
  }

  Key get type {
    return new Key(cboType.value);
  }

  String get size {
    return cboSize.value;
  }

  String get colour {
    return txtColour.value;
  }

  String get material {
    return txtMaterial.value;
  }

  num get weight {
    return numWeight.valueAsNumber;
  }

  void onSend(Event e) {
    if (isFormValid()) {
      disableSubmit(true);
      submitSend();
    }
  }

  submitSend() async {
    final obj = new Clothing(
        code, brand, description, type, size, colour, material, weight);

    HttpRequest req;
    if (objKey.toJson() != "0`0") {
      req = await updateClothing(objKey, obj);
      if (req.status == 200) {
        Toast.success(
            title: "Success!",
            message: req.response,
            position: ToastPos.bottomLeft);
      } else {
        Toast.error(
            title: "Failed!",
            message: req.response,
            position: ToastPos.bottomLeft);
      }
    } else {
      var req = await createClothing(obj);

      if (req.status == 200) {
        final key = req.response;
        objKey = new Key(key);

        new Toast.success(
            title: "Success!",
            message: req.response,
            position: ToastPos.bottomLeft);
        super.form.reset();
      } else {
        new Toast.error(
            title: "Error!",
            message: req.response,
            position: ToastPos.bottomLeft);
      }
    }
  }
}
