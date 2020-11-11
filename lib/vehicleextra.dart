import 'dart:html';
import 'package:mango_stock/bodies/category.dart';
import 'package:mango_ui/formstate.dart';
import 'package:mango_ui/keys.dart';

class VehicleExtra {

  List<String> get extraitems {
    final boxes = querySelectorAll("input[type=checkbox]");
    return boxes.map((e) => e.innerText);
  }

}