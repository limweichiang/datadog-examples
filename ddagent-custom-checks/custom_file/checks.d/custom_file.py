from checks import AgentCheck
from os.path import exists

class FileCheck(AgentCheck):
  def check(self, instance):
    if (exists(instance['file'])):
      # Checks for file, returns 1 if found. Uses custom metrics. Deprecated.
      #self.gauge('file.exists', 1, tags=['file:'+instance['file']])

      # Checks for file, returns OK if found. Monitor using Service Check Monitors.
      self.service_check('file.exists', 0, tags=['file:'+instance['file']])

    else:
      # Checks for file, returns 0 if not found. Uses custom metrics. Deprecated.
      #self.gauge('file.exists', 0, tags=['file:'+instance['file']])

      # Checks for file, returns CRITICAL if not found. Monitor using Service Check Monitors.
      self.service_check('file.exists', 2, tags=['file:'+instance['file']])
