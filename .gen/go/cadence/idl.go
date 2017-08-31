// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Code generated by thriftrw v1.6.0. DO NOT EDIT.
// @generated

package cadence

import (
	"github.com/uber/cadence/.gen/go/shared"
	"go.uber.org/thriftrw/thriftreflect"
)

var ThriftModule = &thriftreflect.ThriftModule{Name: "cadence", Package: "github.com/uber/cadence/.gen/go/cadence", FilePath: "cadence.thrift", SHA1: "ef869985bf0cd5faaf914d4ad0354651c8b6e4b1", Includes: []*thriftreflect.ThriftModule{shared.ThriftModule}, Raw: rawIDL}

const rawIDL = "// Copyright (c) 2017 Uber Technologies, Inc.\n//\n// Permission is hereby granted, free of charge, to any person obtaining a copy\n// of this software and associated documentation files (the \"Software\"), to deal\n// in the Software without restriction, including without limitation the rights\n// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell\n// copies of the Software, and to permit persons to whom the Software is\n// furnished to do so, subject to the following conditions:\n//\n// The above copyright notice and this permission notice shall be included in\n// all copies or substantial portions of the Software.\n//\n// THE SOFTWARE IS PROVIDED \"AS IS\", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR\n// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,\n// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE\n// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER\n// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,\n// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN\n// THE SOFTWARE.\n\ninclude \"shared.thrift\"\n\nnamespace java com.uber.cadence\n\n/**\n* WorkflowService API is exposed to provide support for long running applications.  Application is expected to call\n* StartWorkflowExecution to create an instance for each instance of long running workflow.  Such applications are expected\n* to have a worker which regularly polls for DecisionTask and ActivityTask from the WorkflowService.  For each\n* DecisionTask, application is expected to process the history of events for that session and respond back with next\n* decisions.  For each ActivityTask, application is expected to execute the actual logic for that task and respond back\n* with completion or failure.  Worker is expected to regularly heartbeat while activity task is running.\n**/\nservice WorkflowService {\n  /**\n  * RegisterDomain creates a new domain which can be used as a container for all resources.  Domain is a top level\n  * entity within Cadence, used as a container for all resources like workflow executions, tasklists, etc.  Domain\n  * acts as a sandbox and provides isolation for all resources within the domain.  All resources belongs to exactly one\n  * domain.\n  **/\n  void RegisterDomain(1: shared.RegisterDomainRequest registerRequest)\n    throws (\n      1: shared.BadRequestError badRequestError,\n      2: shared.InternalServiceError internalServiceError,\n      3: shared.DomainAlreadyExistsError domainExistsError,\n    )\n\n  /**\n  * DescribeDomain returns the information and configuration for a registered domain.\n  **/\n  shared.DescribeDomainResponse DescribeDomain(1: shared.DescribeDomainRequest describeRequest)\n    throws (\n      1: shared.BadRequestError badRequestError,\n      2: shared.InternalServiceError internalServiceError,\n      3: shared.EntityNotExistsError entityNotExistError,\n    )\n\n  /**\n  * UpdateDomain is used to update the information and configuration for a registered domain.\n  **/\n  shared.UpdateDomainResponse UpdateDomain(1: shared.UpdateDomainRequest updateRequest)\n      throws (\n        1: shared.BadRequestError badRequestError,\n        2: shared.InternalServiceError internalServiceError,\n        3: shared.EntityNotExistsError entityNotExistError,\n      )\n\n  /**\n  * DeprecateDomain us used to update status of a registered domain to DEPRECATED.  Once the domain is deprecated\n  * it cannot be used to start new workflow executions.  Existing workflow executions will continue to run on\n  * deprecated domains.\n  **/\n  void DeprecateDomain(1: shared.DeprecateDomainRequest deprecateRequest)\n    throws (\n      1: shared.BadRequestError badRequestError,\n      2: shared.InternalServiceError internalServiceError,\n      3: shared.EntityNotExistsError entityNotExistError,\n    )\n\n  /**\n  * StartWorkflowExecution starts a new long running workflow instance.  It will create the instance with\n  * 'WorkflowExecutionStarted' event in history and also schedule the first DecisionTask for the worker to make the\n  * first decision for this instance.  It will return 'WorkflowExecutionAlreadyStartedError', if an instance already\n  * exists with same workflowId.\n  **/\n  shared.StartWorkflowExecutionResponse StartWorkflowExecution(1: shared.StartWorkflowExecutionRequest startRequest)\n    throws (\n      1: shared.BadRequestError badRequestError,\n      2: shared.InternalServiceError internalServiceError,\n      3: shared.WorkflowExecutionAlreadyStartedError sessionAlreadyExistError,\n      4: shared.ServiceBusyError serviceBusyError,\n    )\n\n  /**\n  * Returns the history of specified workflow execution.  It fails with 'EntityNotExistError' if speficied workflow\n  * execution in unknown to the service.\n  **/\n  shared.GetWorkflowExecutionHistoryResponse GetWorkflowExecutionHistory(1: shared.GetWorkflowExecutionHistoryRequest getRequest)\n    throws (\n      1: shared.BadRequestError badRequestError,\n      2: shared.InternalServiceError internalServiceError,\n      3: shared.EntityNotExistsError entityNotExistError,\n      4: shared.ServiceBusyError serviceBusyError,\n    )\n\n  /**\n  * PollForDecisionTask is called by application worker to process DecisionTask from a specific taskList.  A\n  * DecisionTask is dispatched to callers for active workflow executions, with pending decisions.\n  * Application is then expected to call 'RespondDecisionTaskCompleted' API when it is done processing the DecisionTask.\n  * It will also create a 'DecisionTaskStarted' event in the history for that session before handing off DecisionTask to\n  * application worker.\n  **/\n  shared.PollForDecisionTaskResponse PollForDecisionTask(1: shared.PollForDecisionTaskRequest pollRequest)\n    throws (\n      1: shared.BadRequestError badRequestError,\n      2: shared.InternalServiceError internalServiceError,\n      3: shared.ServiceBusyError serviceBusyError,\n    )\n\n  /**\n  * RespondDecisionTaskCompleted is called by application worker to complete a DecisionTask handed as a result of\n  * 'PollForDecisionTask' API call.  Completing a DecisionTask will result in new events for the workflow execution and\n  * potentially new ActivityTask being created for corresponding decisions.  It will also create a DecisionTaskCompleted\n  * event in the history for that session.  Use the 'taskToken' provided as response of PollForDecisionTask API call\n  * for completing the DecisionTask.\n  **/\n  void RespondDecisionTaskCompleted(1: shared.RespondDecisionTaskCompletedRequest completeRequest)\n    throws (\n      1: shared.BadRequestError badRequestError,\n      2: shared.InternalServiceError internalServiceError,\n      3: shared.EntityNotExistsError entityNotExistError,\n    )\n\n  /**\n  * PollForActivityTask is called by application worker to process ActivityTask from a specific taskList.  ActivityTask\n  * is dispatched to callers whenever a ScheduleTask decision is made for a workflow execution.\n  * Application is expected to call 'RespondActivityTaskCompleted' or 'RespondActivityTaskFailed' once it is done\n  * processing the task.\n  * Application also needs to call 'RecordActivityTaskHeartbeat' API within 'heartbeatTimeoutSeconds' interval to\n  * prevent the task from getting timed out.  An event 'ActivityTaskStarted' event is also written to workflow execution\n  * history before the ActivityTask is dispatched to application worker.\n  **/\n  shared.PollForActivityTaskResponse PollForActivityTask(1: shared.PollForActivityTaskRequest pollRequest)\n    throws (\n      1: shared.BadRequestError badRequestError,\n      2: shared.InternalServiceError internalServiceError,\n      3: shared.ServiceBusyError serviceBusyError,\n    )\n\n  /**\n  * RecordActivityTaskHeartbeat is called by application worker while it is processing an ActivityTask.  If worker fails\n  * to heartbeat within 'heartbeatTimeoutSeconds' interval for the ActivityTask, then it will be marked as timedout and\n  * 'ActivityTaskTimedOut' event will be written to the workflow history.  Calling 'RecordActivityTaskHeartbeat' will\n  * fail with 'EntityNotExistsError' in such situations.  Use the 'taskToken' provided as response of\n  * PollForActivityTask API call for heartbeating.\n  **/\n  shared.RecordActivityTaskHeartbeatResponse RecordActivityTaskHeartbeat(1: shared.RecordActivityTaskHeartbeatRequest heartbeatRequest)\n    throws (\n      1: shared.BadRequestError badRequestError,\n      2: shared.InternalServiceError internalServiceError,\n      3: shared.EntityNotExistsError entityNotExistError,\n    )\n\n  /**\n  * RespondActivityTaskCompleted is called by application worker when it is done processing an ActivityTask.  It will\n  * result in a new 'ActivityTaskCompleted' event being written to the workflow history and a new DecisionTask\n  * created for the workflow so new decisions could be made.  Use the 'taskToken' provided as response of\n  * PollForActivityTask API call for completion. It fails with 'EntityNotExistsError' if the taskToken is not valid\n  * anymore due to activity timeout.\n  **/\n  void  RespondActivityTaskCompleted(1: shared.RespondActivityTaskCompletedRequest completeRequest)\n    throws (\n      1: shared.BadRequestError badRequestError,\n      2: shared.InternalServiceError internalServiceError,\n      3: shared.EntityNotExistsError entityNotExistError,\n    )\n\n  /**\n  * RespondActivityTaskFailed is called by application worker when it is done processing an ActivityTask.  It will\n  * result in a new 'ActivityTaskFailed' event being written to the workflow history and a new DecisionTask\n  * created for the workflow instance so new decisions could be made.  Use the 'taskToken' provided as response of\n  * PollForActivityTask API call for completion. It fails with 'EntityNotExistsError' if the taskToken is not valid\n  * anymore due to activity timeout.\n  **/\n  void  RespondActivityTaskFailed(1: shared.RespondActivityTaskFailedRequest failRequest)\n    throws (\n      1: shared.BadRequestError badRequestError,\n      2: shared.InternalServiceError internalServiceError,\n      3: shared.EntityNotExistsError entityNotExistError,\n    )\n\n  /**\n  * RespondActivityTaskCanceled is called by application worker when it is successfully canceled an ActivityTask.  It will\n  * result in a new 'ActivityTaskCanceled' event being written to the workflow history and a new DecisionTask\n  * created for the workflow instance so new decisions could be made.  Use the 'taskToken' provided as response of\n  * PollForActivityTask API call for completion. It fails with 'EntityNotExistsError' if the taskToken is not valid\n  * anymore due to activity timeout.\n  **/\n  void RespondActivityTaskCanceled(1: shared.RespondActivityTaskCanceledRequest canceledRequest)\n    throws (\n      1: shared.BadRequestError badRequestError,\n      2: shared.InternalServiceError internalServiceError,\n      3: shared.EntityNotExistsError entityNotExistError,\n    )\n\n  /**\n  * RequestCancelWorkflowExecution is called by application worker when it wants to request cancellation of a workflow instance.\n  * It will result in a new 'WorkflowExecutionCancelRequested' event being written to the workflow history and a new DecisionTask\n  * created for the workflow instance so new decisions could be made. It fails with 'EntityNotExistsError' if the workflow is not valid\n  * anymore due to completion or doesn't exist.\n  **/\n  void RequestCancelWorkflowExecution(1: shared.RequestCancelWorkflowExecutionRequest cancelRequest)\n      throws (\n        1: shared.BadRequestError badRequestError,\n        2: shared.InternalServiceError internalServiceError,\n        3: shared.EntityNotExistsError entityNotExistError,\n        4: shared.CancellationAlreadyRequestedError cancellationAlreadyRequestedError,\n        5: shared.ServiceBusyError serviceBusyError,\n      )\n\n  /**\n  * SignalWorkflowExecution is used to send a signal event to running workflow execution.  This results in\n  * WorkflowExecutionSignaled event recorded in the history and a decision task being created for the execution.\n  **/\n  void SignalWorkflowExecution(1: shared.SignalWorkflowExecutionRequest signalRequest)\n    throws (\n      1: shared.BadRequestError badRequestError,\n      2: shared.InternalServiceError internalServiceError,\n      3: shared.EntityNotExistsError entityNotExistError,\n      4: shared.ServiceBusyError serviceBusyError,\n    )\n\n  /**\n  * TerminateWorkflowExecution terminates an existing workflow execution by recording WorkflowExecutionTerminated event\n  * in the history and immediately terminating the execution instance.\n  **/\n  void TerminateWorkflowExecution(1: shared.TerminateWorkflowExecutionRequest terminateRequest)\n    throws (\n      1: shared.BadRequestError badRequestError,\n      2: shared.InternalServiceError internalServiceError,\n      3: shared.EntityNotExistsError entityNotExistError,\n      4: shared.ServiceBusyError serviceBusyError,\n    )\n\n  /**\n  * ListOpenWorkflowExecutions is a visibility API to list the open executions in a specific domain.\n  **/\n  shared.ListOpenWorkflowExecutionsResponse ListOpenWorkflowExecutions(1: shared.ListOpenWorkflowExecutionsRequest listRequest)\n    throws (\n      1: shared.BadRequestError badRequestError,\n      2: shared.InternalServiceError internalServiceError,\n      3: shared.EntityNotExistsError entityNotExistError,\n      4: shared.ServiceBusyError serviceBusyError,\n    )\n\n  /**\n  * ListClosedWorkflowExecutions is a visibility API to list the closed executions in a specific domain.\n  **/\n  shared.ListClosedWorkflowExecutionsResponse ListClosedWorkflowExecutions(1: shared.ListClosedWorkflowExecutionsRequest listRequest)\n    throws (\n      1: shared.BadRequestError badRequestError,\n      2: shared.InternalServiceError internalServiceError,\n      3: shared.EntityNotExistsError entityNotExistError,\n      4: shared.ServiceBusyError serviceBusyError,\n    )\n}"
