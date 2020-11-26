import 'package:Admin.APP/serviceform.dart';
import 'package:mango_ui/keys.dart';

void main() {
  final fromParam = Uri.base.queryParameters['from'] ?? "0`0";
  final from = new Key(fromParam);
  new ServiceForm(new Key("0`0"), from);
}
