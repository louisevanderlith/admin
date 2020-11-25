import 'dart:html';

import 'package:dart_toast/dart_toast.dart';
import 'package:mango_ui/formstate.dart';
import 'package:mango_ui/keys.dart';
import 'package:mango_utility/bodies/service.dart';
import 'package:mango_utility/servicesapi.dart';

class ServiceForm extends FormState {
  Key objKey;
  SelectElement cboDays;
  SelectElement cboHours;
  SelectElement cboMinutes;
  LocalDateTimeInputElement txtStartTime;
  TextInputElement txtDescription;
  TextInputElement txtLocation;

  ServiceForm(Key k) : super("#frmService", "#btnSubmit") {
    objKey = k;
    cboDays = querySelector("#cboDays");
    cboHours = querySelector("#cboHours");
    cboMinutes = querySelector("#cboMinutes");
    txtDescription = querySelector("#txtDescription");
    txtStartTime = querySelector("#txtStartTime");
    txtDescription = querySelector("#txtDescription");
    txtLocation = querySelector("#txtLocation");

    querySelector("#btnSubmit").onClick.listen(onSend);
  }

  String get description {
    return txtDescription.value;
  }

  String get location {
    return txtLocation.value;
  }

  Duration get duration {
    final days = num.parse(cboDays.value);
    final hours = num.parse(cboHours.value);
    final mins = num.parse(cboMinutes.value);
    return new Duration(days: days, hours: hours, minutes: mins);
  }

  DateTime get starttime {
    return DateTime.parse(txtStartTime.value);
  }

  void onSend(Event e) {
    if (isFormValid()) {
      disableSubmit(true);
      submitSend();
    }
  }

  submitSend() async {
    final obj = new Service(duration, starttime, location, description);

    HttpRequest req;
    if (objKey.toJson() != "0`0") {
      req = await updateService(objKey, obj);
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
      var req = await createService(obj);

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
