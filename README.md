# Trials using a receipt printer

This was inspired by this post https://www.laurieherault.com/articles/a-thermal-receipt-printer-cured-my-procrastination from [Hacker News](https://news.ycombinator.com/)

I started with python but didn't get the following to work so switched to a go libary which is working well.

Note that it doesn't work on Windows.

```python
from escpos.printer import Usb

""" Seiko Epson Corp. Receipt Printer (EPSON TM-T88III) """
p = Usb(0x04b8, 0x0e28, 0, profile="TM-T20III")
p.text("Hello World\n")
p.image("logo.gif")
p.barcode('4006381333931', 'EAN13', 64, 2, '', '')
p.cut()
```