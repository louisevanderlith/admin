import 'dart:html';

import 'package:dart_toast/dart_toast.dart';
import 'package:mango_house/bodies/property.dart';
import 'package:mango_house/propertiesapi.dart';
import 'package:mango_ui/formstate.dart';
import 'package:mango_ui/keys.dart';

class PropertyForm extends FormState {
  Key objKey;

  NumberInputElement numBedrooms;
  NumberInputElement numBathrooms;
  NumberInputElement numGarages;
  CheckboxInputElement chkPool;
  TextInputElement txtAddress;
  TextInputElement txtSize;
  TextInputElement txtType;

  PropertyForm(Key k) : super("#frmProperty", "#btnSubmit") {
    objKey = k;
    querySelector("#btnSubmit").onClick.listen(onSubmitClick);
  }

  num get bedrooms {
    return numBedrooms.valueAsNumber;
  }

  num get bathrooms {
    return numBathrooms.valueAsNumber;
  }

  num get garages {
    return numGarages.valueAsNumber;
  }

  bool get pool {
    return chkPool.checked;
  }

  String get address {
    return txtAddress.value;
  }

  String get size {
    return txtSize.value;
  }

  String get type {
    return txtType.value;
  }

  void onSubmitClick(MouseEvent e) async {
    if (isFormValid()) {
      disableSubmit(true);

      final obj =
          new Property(bedrooms, bathrooms, garages, pool, address, size, type);

      HttpRequest req;
      if (objKey.toJson() != "0`0") {
        req = await updateProperty(objKey, obj);
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
        req = await createProperty(obj);

        if (req.status == 200) {
          final key = req.response;
          objKey = new Key(key);

          Toast.success(
              title: "Success!",
              message: "Property Saved",
              position: ToastPos.bottomLeft);
        } else {
          Toast.error(
              title: "Failed!",
              message: req.response,
              position: ToastPos.bottomLeft);
        }
      }
    }
  }
}
