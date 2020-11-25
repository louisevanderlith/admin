import 'dart:html';

import 'package:mango_vehicle/bodies/series.dart';

class VehicleSeries {
  TextInputElement txtModel;
  TextInputElement txtManufacturer;
  TextInputElement txtAssemblyPlant;
  NumberInputElement numMonth;
  NumberInputElement numYear;
  TextInputElement txtTrim;

  VehicleSeries() {
    txtModel = querySelector("#txtModel");
    txtManufacturer = querySelector("#txtManufacturer");
    txtAssemblyPlant = querySelector("#txtAssemblyPlant");
    numMonth = querySelector("#numMonth");
    numYear = querySelector("#numYear");
    txtTrim = querySelector("#txtTrim");
  }

  String get model {
    return txtModel.value;
  }

  String get manufacturer {
    return txtManufacturer.value;
  }

  String get assemblyPlant {
    return txtAssemblyPlant.value;
  }

  num get month {
    return numMonth.valueAsNumber;
  }

  num get year {
    return numYear.valueAsNumber;
  }

  String get trim {
    return txtTrim.value;
  }

  Series toDto() {
    return new Series(year, month, model, manufacturer, trim, assemblyPlant);
  }
}
