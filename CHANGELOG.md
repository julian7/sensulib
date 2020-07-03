# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/)
and this project adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).

## [Unreleased]

Added:

* Basic metrics API for emitting data in [OpenTSDB](http://opentsdb.net/) format

## [v0.2.2] - Feb 10, 2020

Fixed:

* SizeToHuman: handle zero bytes

## [v0.2.1] - Dec 13, 2019

Fixed:

* Removed go.sum

## [v0.2.0] - Dec 13, 2019

Changes:

* Upgrade to goshipdone v0.4.0

## [v0.1.3] - Nov 17, 2019

Changes:

* Replace experimental mage library with goshipdone

## [v0.1.2] - Oct 27, 2019

Changes:

* better filenames for mage/asset

Fixes:

* `asset.NewTarget()` never filled ArchiveName, therefore it rendered as empty string for `asset.Target.Summarize()`.

## [v0.1.1] - Oct 27, 2019

Fixes:

* mage/asset package name was still showing "sensuasset", from before the code separation.

## [v0.1.0] - Oct 27, 2019

Initial release.

[Unreleased]: https://github.com/julian7/sensulib
[v0.2.2]: https://github.com/julian7/sensulib/releases/tag/v0.2.2
[v0.2.1]: https://github.com/julian7/sensulib/releases/tag/v0.2.1
[v0.2.0]: https://github.com/julian7/sensulib/releases/tag/v0.2.0
[v0.1.3]: https://github.com/julian7/sensulib/releases/tag/v0.1.3
[v0.1.2]: https://github.com/julian7/sensulib/releases/tag/v0.1.2
[v0.1.1]: https://github.com/julian7/sensulib/releases/tag/v0.1.1
[v0.1.0]: https://github.com/julian7/sensulib/releases/tag/v0.1.0
