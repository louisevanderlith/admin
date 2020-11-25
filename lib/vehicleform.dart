import 'dart:html';

import 'package:dart_toast/dart_toast.dart';
import 'package:mango_ui/formstate.dart';
import 'package:mango_ui/keys.dart';
import 'package:mango_vehicle/bodies/vehicle.dart';
import 'package:mango_vehicle/vehicleapi.dart';

import 'vehicleengine.dart';
import 'vehicleextra.dart';
import 'vehiclegearbox.dart';
import 'vehicleinfo.dart';
import 'vehicleseries.dart';

class VehicleForm extends FormState {
  Key objKey;

  VehicleInfo info;
  VehicleSeries series;
  VehicleEngine engine;
  VehicleGearbox gearbox;
  VehicleExtra extra;

  VehicleForm(Key k) : super("#frmVehicle", "#btnSubmit") {
    objKey = k;

    querySelector("#btnSubmit").onClick.listen(onSubmitClick);
  }

  void onSubmitClick(MouseEvent e) async {
    if (isFormValid()) {
      disableSubmit(true);

      final obj = new Vehicle(
          info.vinKey,
          info.fullVIN,
          series.toDto(),
          info.colour,
          info.paintNo,
          engine.toDto(),
          gearbox.toDto(),
          info.bodyStyle,
          info.doors,
          extra.extraitems,
          info.spare,
          info.service,
          info.condition,
          info.issues,
          info.mileage);

      HttpRequest req;
      if (objKey.toJson() != "0`0") {
        req = await updateInfo(objKey, obj);
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
        req = await submitVehicle(obj);

        if (req.status == 200) {
          final key = req.response;
          objKey = new Key(key);

          Toast.success(
              title: "Success!",
              message: "Vehicle Saved",
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
